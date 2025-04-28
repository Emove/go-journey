package code

import "fmt"

func IfExistsTest() {
	scoreMap := make(map[string]int)
	scoreMap["zhangsan"] = 100
	scoreMap["lisi"] = 200
	v, ok := scoreMap["zhangsan"]
	if ok {
		fmt.Println(v)
	} else {
		fmt.Println("no exists")
	}
}

func ForEachMapTest() {
	scoreMap := make(map[string]int)
	scoreMap["zhangsan"] = 90
	scoreMap["lisi"] = 100
	scoreMap["wangwu"] = 50
	for key, value := range scoreMap {
		fmt.Println(key, value)
	}
}

func DeleteFormMapTest() {
	scoreMap := make(map[string]int)
	scoreMap["zhangsan"] = 90
	scoreMap["lisi"] = 100
	scoreMap["wangwu"] = 50
	delete(scoreMap, "lisi")
	delete(scoreMap, "zhangsan")
	delete(scoreMap, "wangwu")
	for key, value := range scoreMap {
		fmt.Println(key, value)
	}
}

func MapSliceTest() {
	var mapSlice = make([]map[string]string, 3)
	for index, value := range mapSlice {
		fmt.Printf("index : %v value : %v \n", index, value)
	}
	mapSlice[0] = make(map[string]string, 10)
	mapSlice[0]["name"] = "zhangsan"
	mapSlice[0]["password"] = "123456"
	mapSlice[0]["address"] = "shenzhen"

	mapSlice[1] = make(map[string]string, 10)
	mapSlice[1]["name"] = "lisi"
	mapSlice[1]["password"] = "654321"
	mapSlice[1]["address"] = "guangzhou"
	for index, value := range mapSlice {
		fmt.Printf("index: %v value:%v \n", index, value)
	}
}

func SliceMapTest() {
	var sliceMap = make(map[string][]string, 3)
	fmt.Println(sliceMap)
	fmt.Println("after init")
	key := "中国"
	value, ok := sliceMap[key]
	if !ok {
		value = make([]string, 0, 2)
	}
	value = append(value, "北京", "上海")
	sliceMap[key] = value
	fmt.Println(sliceMap)
}

func LenOfMap() {
	m := make(map[int]struct{})
	for i := 0; i < 3; i++ {
		m[i] = struct{}{}
	}

	fmt.Println(len(m))
}

func ReadMapK() {
	m := make(map[int][]int)
	for _, v := range m[0] {
		fmt.Println(v)
	}

}
