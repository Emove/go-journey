package oop

import "fmt"

func Assignment2InterfaceByPassThrough() {
	var face interface{}
	assignment(face)
	fmt.Printf("%#v", face)
}

func assignment(face interface{}) {
	face = string("hello world")
}
