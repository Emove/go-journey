package main

import (
	"context"
	"log"
	"net"
	"sync"
	"time"
)

type ConnPair struct {
	client net.Conn
	server net.Conn
}

func ConcurrencyWrite() {
	pair, err := prepare()
	if err != nil {
		log.Fatal(err.Error())
	}

	content := "hello world"

	ctx, cancelFunc := context.WithCancel(context.Background())
	go func() {
		for {
			select {
			case <-ctx.Done():
				return
			default:
				buf := make([]byte, len(content))
				_, err := pair.server.Read(buf)
				if err != nil {
					log.Fatal(err.Error())
				}
				if string(buf) != content {
					log.Fatalf("concurrency write data error, read content: %s", string(buf))
				} else {
					log.Println(string(buf))
				}
			}
		}
	}()

	cnt := 10
	wg := sync.WaitGroup{}
	wg.Add(cnt)
	for i := 0; i < cnt; i++ {
		go func() {
			wg.Wait()
			_, err := pair.client.Write([]byte(content))
			if err != nil {
				log.Fatal(err.Error())
			}
		}()
		wg.Done()
	}
	time.Sleep(1 * time.Second)
	cancelFunc()
}

func prepare() (pair *ConnPair, err error) {
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

	pair = &ConnPair{}
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
