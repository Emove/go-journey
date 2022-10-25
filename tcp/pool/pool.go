package pool

import (
	"errors"
	"sync"
	"sync/atomic"
	"time"
)

// Concrete https://github.com/emove/connpool
type Pool interface {
	Get(addr ...string) (interface{}, error)
	Put(interface{})
	Discard(interface{})
	Close()
}

type (
	// Dialer is used to create new connections.
	Dialer = func(addr string) (interface{}, error)
	// Closer is a connection closed hook.
	Closer = func(interface{})
)

var (
	ErrNotUsableAddr      = errors.New("not usable addr")
	ErrPoolClosed         = errors.New("pool closed")
	ErrGetConnWaitTimeout = errors.New("get connection wait timeout")
	ErrIllegalAddress     = errors.New("address not register")
)

const (
	working = iota
	closed
)

type pool struct {
	mu          sync.RWMutex
	addrs       []string
	idx         int
	peers       map[string]*peer
	closed      int
	done        chan struct{} // the signal to shutdown the process
	dialer      Dialer
	closer      Closer
	reqs        sync.Map
	freeChan    chan *idleConn
	inmu        sync.Mutex
	inuse       map[interface{}]*peer
	coreNums    int
	maxNums     int
	maxIdleTime time.Duration
	lazy        bool
}

func (p *pool) Get(addr ...string) (conn interface{}, err error) {
	p.mu.RLock()
	defer p.mu.RUnlock()
	if p.closed == closed {
		return nil, ErrPoolClosed
	}
	var pr *peer
	if len(addr) > 0 {
		if p.peers[addr[0]] == nil {
			p.mu.RUnlock()
			return nil, ErrIllegalAddress
		}
	} else {
		pr, err = p.roundRobin()
		if err != nil {
			return nil, err
		}
	}

	req := newRequest(pr, p.maxIdleTime)
	v, _ := p.reqs.Load(pr)
	ch := v.(chan *request)
	ch <- req

	select {
	case conn = <-req.conn:
	case err = <-req.err:
	}

	if err != nil {
		if err != ErrGetConnWaitTimeout {
			// Although request return due to timeout, it still
			// in the peer's waiting channel, so it should be
			// reused by the peer's process
			reuseRequest(req)
		}
		return
	}

	if _, ok := conn.(Connection); ok {
		return
	}

	p.inmu.Lock()
	p.inuse[conn] = pr
	p.inmu.Unlock()
	return
}

// Put puts the connection to pool
func (p *pool) Put(conn interface{}) {
	if conn == nil {
		return
	}

	p.mu.RLock()
	defer p.mu.RUnlock()
	if p.closed == closed {
		return
	}

	pr, ok := (*peer)(nil), false
	if con, ok := conn.(Connection); ok {
		v := con.Value(&ctxPeerKey{})
		if v != nil {
			pr = v.(*peer)
		}
	}
	if pr == nil {
		p.inmu.Lock()
		if pr, ok = p.inuse[conn]; !ok {
			p.inmu.Unlock()
			return
		}
		delete(p.inuse, conn)
		p.inmu.Unlock()
	}
	pr.idles <- newIdleConn(pr, conn)
}

func (p *pool) Discard(conn interface{}) {
	if conn == nil {
		return
	}

	p.mu.RLock()
	defer p.mu.RUnlock()

	pr, ok := (*peer)(nil), false
	if con, ok := conn.(Connection); ok {
		v := con.Value(&ctxPeerKey{})
		if v != nil {
			pr = v.(*peer)
		}
	}
	if pr == nil {
		p.inmu.Lock()
		if pr, ok = p.inuse[conn]; !ok {
			p.inmu.Unlock()
			return
		}
		delete(p.inuse, conn)
		p.inmu.Unlock()
	}

	ic := newIdleConn(pr, conn)
	if p.closed == closed {
		p.free(ic)
	} else {
		p.freeChan <- ic
	}
}

func (p *pool) Close() {
	p.mu.Lock()
	defer p.mu.Unlock()

	p.closed = closed
	// shutdown the process loop
	close(p.done)
	p.clear()
}

func (p *pool) process() {
	for {
		select {
		case <-p.done:
			return
		case ic := <-p.freeChan:
			p.free(ic)
		default:
			p.foreachPeerRequest()
		}
	}
}

func (p *pool) foreachPeerRequest() {
	p.reqs.Range(func(key, value interface{}) bool {
		pr, ch := key.(*peer), value.(chan *request)
		loops := 0
	requestLoop:
		for len(ch) > 0 {
			loops++
			if loops > 60 {
				// Handle others request channel
				return true
			}
			ic := (*idleConn)(nil)
			for len(pr.idles) > 0 {
				// Select a usable idle connection
				ic = <-pr.idles
				if p.maxIdleTime > 0 && time.Since(ic.idleTime) > p.maxIdleTime {
					p.freeChan <- ic
					continue
				}
				break
			}

			if ic == nil {
				p.dial(pr)
				break
			}

			for len(ch) > 0 {
				req := <-ch
				if !atomic.CompareAndSwapInt32(&req.state, waiting, finished) {
					// The request timeout
					reuseRequest(req)
					continue
				}

				req.conn <- ic.conn
				reuseIdleConn(ic)
				goto requestLoop
			}

			// All request timeout, return the idle connection to channel
			pr.idles <- ic
		}
		return true
	})
}

func (p *pool) free(ic *idleConn) {
	p.closer(ic.conn)
	pr := ic.pr
	reuseIdleConn(ic)
	pr.am.Lock()
	pr.actives--
	if !p.lazy && pr.actives < p.coreNums {
		pr.actives++
		pr.am.Unlock()
		p.dial(pr)
	}
	pr.am.Unlock()
}

func (p *pool) dial(pr *peer) {
	p.mu.RLock()
	if p.closed == closed {
		pr.am.Lock()
		pr.actives--
		pr.am.Unlock()
		return
	}
	p.mu.RUnlock()
	conn, err := p.dialer(pr.addr)
	if err != nil {
		pr.am.Lock()
		pr.actives--
		pr.am.Unlock()
		return
	}
	if con, ok := conn.(Connection); ok {
		con.WithValue(&ctxPeerKey{}, pr)
	}

	ic := newIdleConn(pr, conn)
	p.mu.RLock()
	p.mu.RUnlock()
	if p.closed == closed {
		p.free(ic)
	} else {
		pr.idles <- ic
	}
}

func (p *pool) clear() {
	// Reject all of the waiting requests
	p.clearReqs()

	// Clear free channel
	for ic := range p.freeChan {
		p.free(ic)
	}

	// Clear all peers
	for _, pr := range p.peers {
		for ic := range pr.idles {
			p.free(ic)
		}
	}
}

func (p *pool) clearReqs() {
	p.reqs.Range(func(key, value interface{}) bool {
		ch := value.(chan *request)
		for len(ch) > 0 {
			req := <-ch
			if !atomic.CompareAndSwapInt32(&req.state, waiting, finished) {
				// The request timeout
				reuseRequest(req)
				continue
			}
			req.err <- ErrPoolClosed
		}
		close(ch)
		return true
	})
}

func (p *pool) roundRobin() (*peer, error) {
	if len(p.addrs) < 1 {
		if p.closed == closed {
			return nil, ErrPoolClosed
		}
		return nil, ErrNotUsableAddr
	}

	if p.idx >= len(p.addrs) {
		p.idx = 0
	}
	addr := p.addrs[p.idx]
	p.idx++
	return p.peers[addr], nil
}
