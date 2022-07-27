package main

import (
	"fmt"
	"net"
)

const defaultProtocol = "tcp4"

type Connection struct {
	ip       string
	port     string
	protocol string
	con      *net.TCPConn
	listener *net.TCPListener
}

func NewConnection(ip, port string) Connection {
	return Connection{
		ip:       ip,
		port:     port,
		protocol: defaultProtocol,
	}
}

func (conn *Connection) Open() {
	address := conn.GetTargetAddress()
	addr, err := net.ResolveTCPAddr(conn.protocol, address)
	if err != nil {
		fmt.Printf("resolve %s failed\n", address)
	}
	tcpConn, err := net.DialTCP(conn.protocol, nil, addr)
	if err != nil {
		fmt.Printf("connect to %s failed\n", address)
	}
	fmt.Printf("connect to %s success\n", address)
	conn.con = tcpConn
}

func (conn *Connection) Listen() {
	address := conn.GetTargetAddress()
	addr, err := net.ResolveTCPAddr(conn.protocol, address)
	if err != nil {
		fmt.Printf("resolve %s failed\n", address)
		return
	}
	listen, err := net.ListenTCP(conn.protocol, addr)
	if err != nil {
		fmt.Printf("listen to %s failed\n", address)
		return
	}
	con, err := listen.AcceptTCP()
	if err != nil {
		fmt.Println("accept a connection failed")
		return
	}
	fmt.Println("accept a connection from ", con.RemoteAddr().String())
	conn.listener = listen
	conn.con = con
}

func (conn *Connection) CloseConnection() {
	_ = conn.con.Close()
}

func (conn *Connection) CloseListener() {
	_ = conn.listener.Close()
}

func (conn *Connection) Write(data string) {
	buf := []byte(data)
	_, err := conn.con.Write(buf)
	if err != nil {
		fmt.Println("write data to channel failed")
	}
}

func (conn *Connection) Read() string {
	buf := make([]byte, 1024)
	_, err := conn.con.Read(buf)
	if err != nil {
		fmt.Println("read channel data failed")
		return ""
	}
	return string(buf)
}

func (conn *Connection) GetTargetAddress() string {
	return fmt.Sprintf("%s:%s", conn.ip, conn.port)
}
