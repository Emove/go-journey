package oop

import (
	"fmt"
)

type Person struct {
	name   string
	gender string
	age    int8
}

type Student struct {
	Person
	*Class
	id   int
	addr string
	name string
}

type Class struct {
	name string
}

func NewStudentTest() {
	stu := Student{
		Person: Person{
			name:   "Emove",
			gender: "male",
			age:    20},
		Class: &Class{
			name: "6班",
		},
		id:   1,
		addr: "shenzhen",
		name: "Emove",
	}
	fmt.Printf("%#v", stu)
}
