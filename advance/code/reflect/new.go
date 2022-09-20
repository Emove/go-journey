package reflect

import (
	"bytes"
	"encoding/json"
	"fmt"
	"reflect"
)

type Class struct {
	Name       string                 `json:"Name"`
	Properties map[string]interface{} `json:"Properties"`
	unexported string                 `json:"unexported"`
}

func NewByType() {
	t := reflect.New(reflect.TypeOf(Class{}))
	j := `{"Name":"Emove","Properties":{"age":"20","gender":"male"}, "unexported": "1"}`
	decoder := json.NewDecoder(bytes.NewBuffer([]byte(j)))
	err := decoder.Decode(t.Interface())
	if err != nil {
		fmt.Printf("%v\n", err)
	}
	fmt.Printf("%v", t)
}

func ByJSON() {
	j := `{"Name":"Emove","Properties":{"age":"20","gender":"male"}, "unexported": "1"}`
	decoder := json.NewDecoder(bytes.NewBuffer([]byte(j)))
	t := &Class{}
	_ = decoder.Decode(t)
	//if err != nil {
	//	fmt.Printf("%v\n", err)
	//}
	//fmt.Printf("%v", t)
}

func Print() {
	c := Class{
		Name:       "Emove",
		Properties: map[string]interface{}{"age": "20", "gender": "male"},
		unexported: "1",
	}
	buf := &bytes.Buffer{}
	encoder := json.NewEncoder(buf)
	err := encoder.Encode(c)
	if err != nil {
		fmt.Printf("%v", err)
		return
	}
	fmt.Printf("%s", buf.String())
}
