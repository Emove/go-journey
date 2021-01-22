package code

import (
	"fmt"
	"reflect"
)

type User struct {
	Name string
	Age  int
	Id   string
}

type Unexported struct {
	id   string
	name string
}

func (target *Unexported) SetID(id string) {
	target.id = id
}

func (u *User) SayHello() {
	fmt.Println("I'm " + u.Name + ", Id is " + u.Id + ". Nice to meet you! ")
}

func TestReflect() {
	tonydon := &User{"TangXiaodong", 100, "0000123"}
	object := reflect.ValueOf(tonydon)
	myref := object.Elem()
	typeOfType := myref.Type()
	for i := 0; i < myref.NumField(); i++ {
		field := myref.Field(i)
		fmt.Printf("%d. %s %s = %v \n", i, typeOfType.Field(i).Name, field.Type(), field.Interface())
	}
	tonydon.SayHello()
	v := object.MethodByName("SayHello")
	v.Call([]reflect.Value{})
}

func SetValue() {
	target := &Unexported{}
	//typeOf := reflect.TypeOf(target)
	//idField, exist := typeOf.FieldByName("id")
	//if exist {
	//	fmt.Println(idField.Type.String())
	//}
	elem := reflect.ValueOf(target).Elem()
	name := elem.FieldByName("id")
	//fmt.Println("id：", )
	fmt.Println("id type:", name.Kind().String())
	valueOf := reflect.ValueOf(&target).Elem()
	method := valueOf.MethodByName("SetID")
	idValue := make([]reflect.Value, 1)
	idValue[0] = reflect.ValueOf("1")
	method.Call(idValue)
	fmt.Printf("%#v", target)
}
