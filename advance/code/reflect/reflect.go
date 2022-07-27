package reflect

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
	v.Call(nil)
}

func SetValue() {
	target := Unexported{}
	//typeOf := reflect.TypeOf(target)
	//idField, exist := typeOf.FieldByName("id")
	//if exist {
	//	fmt.Println(idField.Type.String())
	//}
	targetValue := reflect.ValueOf(target)
	fmt.Printf("%v\n", targetValue)
	//targetType := reflect.TypeOf(target)
	fmt.Printf("%d\n", targetValue.NumField())
	name := targetValue.FieldByName("id")
	//fmt.Println("id：", )
	fmt.Println("id type:", name.Kind().String())

	//idValue := make([]reflect.Value, 1)
	//idValue[0] = reflect.ValueOf("1")

	fmt.Printf("%v\n", targetValue.Field(0).CanSet())
	//targetValue.Field(0).Set(reflect.ValueOf("1"))

	method := reflect.ValueOf(&target).MethodByName("SetID")
	method.Call([]reflect.Value{reflect.ValueOf("1")})
	fmt.Printf("%#v\n", target)
}
