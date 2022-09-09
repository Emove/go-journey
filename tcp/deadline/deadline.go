package deadline

import (
	"fmt"
	"log"
	"net"
	"sync"
	"time"
)

type connPair struct {
	client net.Conn
	server net.Conn
}

func prepare() (pair *connPair, err error) {
	network, addr := "tcp", ":8080"
	listen, err := net.Listen(network, addr)
	if err != nil {
		return nil, err
	}

	type clientCon struct {
		client net.Conn
		err    error
	}

	clientChann := make(chan *clientCon)
	go func() {
		time.Sleep(1 * time.Second)
		client, err1 := net.Dial(network, addr)
		clientChann <- &clientCon{client: client, err: err1}
	}()

	pair = &connPair{}
	defer func() {
		if err != nil {
			if pair.server != nil {
				_ = pair.server.Close()
			}
			if pair.client != nil {
				_ = pair.client.Close()
			}
		}
	}()

	for {
		var con net.Conn
		con, err = listen.Accept()
		if err != nil {
			return
		}
		pair.server = con
		break
	}

	select {
	case cc := <-clientChann:
		if cc.err != nil {
			return nil, err
		}
		pair.client = cc.client
		return
	}
}

type fn func(pair *connPair)

func do(f fn) {
	pair, err := prepare()
	if err != nil {
		log.Fatalln(err)
	}

	defer func() {
		if e := recover(); e != nil {
			if pair.client != nil {
				_ = pair.client.Close()
			}

			if pair.server != nil {
				_ = pair.server.Close()
			}
		}
	}()

	f(pair)

	if pair.client != nil {
		_ = pair.client.Close()
	}

	if pair.server != nil {
		_ = pair.server.Close()
	}
}

func SetDeadline() {
	do(func(pair *connPair) {

		client := pair.client

		msg := []byte("hello world")
		buff := make([]byte, len(msg))

		wg := &sync.WaitGroup{}
		wg.Add(1)
		go func() {
			err := client.SetReadDeadline(time.Now().Add(time.Second))
			if err != nil {
				fmt.Printf("set deadline err: %v\n", err)
			}
			wg.Done()
			if _, err := client.Read(buff); err != nil {
				if err != nil {
					fmt.Printf("read err: %v\n", err)
				}
			}
			fmt.Printf("read: %s\n", string(buff))
		}()

		wg.Wait()
		time.AfterFunc(time.Second/2-time.Millisecond*3, func() {
			if _, err := pair.server.Write(msg); err != nil {
				fmt.Printf("write err: %v\n", err)
			}
		})

	})
}
