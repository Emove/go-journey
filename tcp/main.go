package main

import (
	"fmt"
	"time"
)

func main() {
	go server()
	client()
}

func client() {
	con := NewConnection("127.0.0.1", "8888")
	con.Open()
	defer con.CloseConnection()
	for i := 0; i < 4; i++ {
		con.Write("heartbeat...")
		response := con.Read()
		fmt.Println("client receive: ", response)
		time.Sleep(time.Second * 5)
	}
}

func server() {
	con := NewConnection("127.0.0.1", "8888")
	con.Listen()
	defer con.CloseListener()
	for {
		msg := con.Read()
		if "" != msg {
			fmt.Println("server receive: ", msg)
			con.Write("ok!")
		}
	}
}
