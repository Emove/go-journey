package code

import "fmt"

type person struct {
	name string
	city string
	age int8
}

func PrintPersonInstance() {
	var instance person
	instance.name = "emove"
	instance.city = "shenzhen"
	instance.age = 20
	fmt.Println(instance)
	fmt.Printf("person=%#v\n", instance)
}

func CreatePersonInstanceByNew() {
	var per = new(person)
	per.name = "Emove"
	per.age = 20
	per.city = "sz"
	fmt.Printf("person=%#v\n", per)
}