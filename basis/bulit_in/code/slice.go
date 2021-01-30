package code

import "fmt"

func CreateSlice() {
	// 1、声明切片
	var s1 []int

	if s1 == nil {
		fmt.Println("s1 is nil")
	} else {
		fmt.Println("s1 is not nil")
	}

	// 2、：=
	s2 := []int{}

	// 3、make()
	var s3 []int = make([]int, 0)
	fmt.Println(s1, s2, s3)

	// 4、初始化赋值
	var s4 []int = make([]int, 0, 0)
	fmt.Println(s4)

	s5 := []int{1, 2, 3}
	fmt.Println(s5)

	// 5、从数组切片
	arr := [5]int{1, 2, 3, 4, 5}
	//var s6[] int
	// 包前不包后
	//s6 = arr[1:4]
	fmt.Println(arr[:1])
}

var arr = [...]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
var slice0 []int = arr[2:8]
var slice1 []int = arr[:6]
var slice2 []int = arr[5:]
var slice3 []int = arr[:len(arr)-1]
var slice4 = arr[:len(arr)-1]

func PrintVariable() {
	fmt.Printf("全局变量： arr %v\n", arr)
	fmt.Printf("全局变量： slice0 %v\n", slice0)
	fmt.Printf("全局变量： slice1 %v\n", slice1)
	fmt.Printf("全局变量： slice2 %v\n", slice2)
	fmt.Printf("全局变量： slice3 %v\n", slice3)
	fmt.Printf("全局变量： slice4 %v\n", slice4)

	fmt.Println("-----------------------------------")
	arr := [...]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	slice0 := arr[2:8]
	slice1 := arr[:6]
	slice2 := arr[5:]
	slice3 := arr[:]
	slice4 := arr[0 : len(arr)-1]

	fmt.Printf("局部变量 arr %v\n", arr)
	fmt.Printf("局部变量 slice0 %v\n", slice0)
	fmt.Printf("局部变量 slice1 %v\n", slice1)
	fmt.Printf("局部变量 slice2 %v\n", slice2)
	fmt.Printf("局部变量 slice3 %v\n", slice3)
	fmt.Printf("局部变量 slice4 %v\n", slice4)
}

func PrintCreateByMake() {
	// var slice []type = make([]type, len)
	// slice  := make([]type, len)
	// slice  := make([]type, len, cap)
	var slice01 []int = make([]int, 10)
	var slice02 = make([]int, 10)
	var slice03 = make([]int, 10, 10)
	fmt.Printf("make局部变量slice01： %v \n", slice01)
	fmt.Printf("make局部变量slice02： %v \n", slice02)
	fmt.Printf("make局部变量slice03： %v \n", slice03)
}

func PrintModify() {
	data := [...]int{0, 1, 2, 3, 4, 5}
	s := data[2:4]
	s[0] += 100
	s[1] += 200

	fmt.Println(data)
	fmt.Println(s)

	s1 := []int{0, 1, 2, 3}
	p := &s1[2]
	*p += 100
	fmt.Println(s1)
}

func PrintInit() {
	s1 := []int{0, 1, 2, 3, 8: 100}
	fmt.Println(s1)

	// 指定len和cap
	s2 := make([]int, 5, 6)
	fmt.Println(s2, len(s2), cap(s2))

	// 省略cap，相当与len = cap
	s3 := make([]int, 6)
	fmt.Println(s3, len(s3), cap(s3))
}

func PrintTwoDimensional() {
	s := [][]int{
		[]int{0, 1, 2},
		[]int{3, 4, 5, 6},
		[]int{7, 8, 9},
	}
	fmt.Println(s)
}

func AppendTest() {
	var a = []int{1, 2, 3}
	fmt.Printf("slice a : %v\n", a)

	var b = []int{4, 5, 6}

	c := append(a, b...)
	fmt.Printf("slice c : %v\n", c)

	d := append(c, 7)
	fmt.Printf("slice d : %v\n", d)

	s1 := make([]int, 0, 5)
	fmt.Printf("slice s1 : %v\n", s1)

	s2 := append(s1, 1, 2)
	fmt.Printf("slice s2 : %v\n", s2)
}

func AppendOutOfCapTest() {
	data := [...]int{0, 1, 2, 3, 4, 5, 10: 10}
	s := data[:2:3]
	fmt.Printf("slice s: %v\n", s)
	// 一次append两个值，超出cap限制
	s = append(s, 100, 200)
	fmt.Println(s, data)
	fmt.Println(&s[0], &data[0])
}

// 2的倍数
func SliceAssignRuleTest() {
	s := make([]int, 0, 1)
	c := cap(s)

	for i := 0; i < 50; i++ {
		s = append(s, i)
		if n := cap(s); n > c {
			fmt.Printf("cap %d -> %d\n", c, n)
			c = n
		}
	}
}

// 函数copy在两个slice间复制数据，复制长度以len小的为准。
// 两个slice可指向同一个底层数组，允许元素区间重叠
func SliceCopyTest() {
	s1 := []int{1, 2, 3, 4, 5}
	fmt.Printf("slice s1: %v\n", s1)
	s2 := make([]int, 10)
	fmt.Printf("slice s2: %v\n", s2)

	copy(s2, s1)
	fmt.Printf("copied slice s1: %v\n", s1)
	fmt.Printf("copied slice s2: %v\n", s2)
}

func ForEachSliceTest() {
	data := [...]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	slice := data[:]

	for index, value := range slice {
		fmt.Printf("index: %v, value: %v \n", index, value)
	}
}

func SliceResizeTest() {
	var a = []int{1, 2, 3, 4, 5}
	fmt.Printf("slice a: %v, len: %v\n", a, len(a))
	b := a[1:2]
	fmt.Printf("slice b: %v, len: %v\n", b, len(b))
	c := b[0:3]
	fmt.Printf("slice c: %v, len: %v\n", c, len(c))

	fmt.Println(&a, &b, &c)
	fmt.Println(&a[1], &b[0], &c[0])
}

func SliceStringTest() {
	str := "hello world"
	s1 := str[0:5]
	fmt.Println(s1)

	s2 := str[6:]
	fmt.Println(s2)
}

func SliceChangeStringTest1() {
	str := "Hello World"
	s := []byte(str)
	s[6] = 'G'
	s = s[:8]
	s = append(s, '!')
	str = string(s)
	fmt.Println(str)
}

func SliceChangeStringTest2() {
	str := "你好，世界！"
	s := []rune(str)
	s[3] = '够'
	s[4] = '浪'
	s = s[0:6]
	fmt.Println(string(s))
}

func AppendAndForeachTest() {
	nums1 := make([]int, 0)
	for i := 0; i < 5; i++ {
		nums1 = append(nums1, i)
	}

	for i := range nums1 {
		fmt.Println(nums1[i])
	}

	nums2 := make([]int, 10, 10)
	for i := 0; i < 5; i++ {
		nums2 = append(nums2, i)
	}

	for i := range nums2 {
		fmt.Println(nums2[i])
	}

}
