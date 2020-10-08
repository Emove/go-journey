/*
 * @author: Emove
 * @date: Do not edit
 */
package code

import "fmt"

type person struct {
	name string
	city string
	age  int8
}

type User struct {
	name    string
	gender  string
	address Address
}

type Address struct {
	province, city string
}

type Animal struct {
	class string
}

type Dog struct {
	name    string
	*Animal // 通过嵌套匿名结构体实现继承
}

func PrintPersonInstance() {
	var instance person
	instance.name = "emove"
	instance.city = "shenzhen"
	instance.age = 20
	fmt.Println(instance)
	fmt.Printf("person=%#v\n", instance)
}

func CreatePersonInstance() {
	per := &person{}
	per.name = "Emove"
	per.age = 20
	per.city = "sz"

	per1 := &person{
		name: "emove",
		city: "st",
		age:  18,
	}

	per2 := &person{
		"e", "sawtou", 19,
	}
	fmt.Printf("person=%#v\n", per)
	fmt.Printf("person1=%#v\n", per1)
	fmt.Printf("person2=%#v\n", per2)
}

func CreatePersonInstanceByNew() {
	var per = new(person)
	per.name = "Emove"
	per.age = 20
	per.city = "sz"
	fmt.Printf("person=%#v\n", per)
}

func PutPersonInfo2Map() {
	m := make(map[string]*person)
	pers := [...]person{
		{name: "Emove", age: 20, city: "sz"},
		{name: "Ez", age: 22, city: "hz"},
		{name: "ff", age: 22, city: "gz"},
	}

	for _, per := range pers {
		fmt.Printf("person = %p\n", &per)
		m[per.name] = &per
	}

	for key, value := range m {
		fmt.Println(key, " ---> ", value.name)
	}
	m1 := make(map[string]*person)
	for i := 0; i < len(pers); i++ {
		per := pers[i]
		m1[per.name] = &per
	}

	for key, value := range m1 {
		fmt.Println(key, " ---> ", value.name)
	}
}

func Construct(name, city string, age int8) *person {
	return &person{
		name: name,
		city: city,
		age:  age,
	}
}

func (p person) Dream() {
	fmt.Printf("%s的梦想是学好go language\n", p.name)
}

func (per *person) SetAge(age int8) {
	per.age = age
}

func NewNestedStruct() {
	user := &User{
		name:   "Emove",
		gender: "male",
		address: Address{
			province: "guangdong",
			city:     "shenzhen",
		},
	}
	fmt.Printf("user = %#v\n", user)
}

func (dog *Dog) Yelp() {
	fmt.Printf("%s ---> 汪汪汪...\n", dog.name)
}

func (animal *Animal) Move() {
	fmt.Printf("%s ---> moving\n", animal.class)
}

func NewDogStruct() *Dog {
	return &Dog{
		name: "dahuang",
		Animal: &Animal{
			class: "dog",
		},
	}
}
