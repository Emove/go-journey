package pool

import (
	"sync"
	"time"
)

type ctxPeerKey struct{}

type peer struct {
	p       *pool
	addr    string
	idles   chan *idleConn
	am      sync.Mutex
	actives int
}

type idleConn struct {
	pr       *peer
	conn     interface{}
	idleTime time.Time
}

var idleConnPool sync.Pool

func init() {
	idleConnPool = sync.Pool{New: func() interface{} {
		return &idleConn{}
	}}
}

func newIdleConn(pr *peer, conn interface{}) *idleConn {
	ic := idleConnPool.Get().(*idleConn)
	ic.pr = pr
	ic.conn = conn
	ic.idleTime = time.Now()
	return ic
}

func reuseIdleConn(ic *idleConn) {
	if ic == nil {
		return
	}
	ic.idleTime = time.Time{}
	ic.conn = nil
	ic.pr = nil
	idleConnPool.Put(ic)
}
