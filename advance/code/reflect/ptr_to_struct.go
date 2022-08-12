package reflect

import (
	"fmt"
	"reflect"
)

type MyStruct struct {
	A int
	B []int
}

func ConvertPtr2struct() {
	s := MyStruct{
		A: 1,
		B: []int{2, 3},
	}
	p := &MyStruct{
		A: 4,
		B: []int{5, 6},
	}

	if reflect.ValueOf(s).Type().Kind() == reflect.Struct {
		fmt.Println("b is a struct")
	}

	if reflect.ValueOf(p).Type().Kind() == reflect.Ptr {
		fmt.Println("p is a pointer")
	}

	field := reflect.ValueOf(p).Field(0)
	fmt.Println(field.Int())

	indirect := reflect.Indirect(reflect.ValueOf(p))
	if indirect.Type().Kind() == reflect.Struct {
		fmt.Println("p has bean convert to struct")
	}
	field = indirect.Field(0)
	fmt.Println(field.Int())
}
