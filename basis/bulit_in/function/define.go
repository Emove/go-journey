/*
 * @author: Emove
 * @date: Do not edit
 */
/*
 * @author: Emove
 * @date: Do not edit
 */
package function

import (
	"fmt"
	"math"
)

// 定义函数类型
type FormatFunc func(s string, x, y int) string

func test(fn func() int) int {
	return fn()
}

func format(fn FormatFunc, s string, x, y int) string {
	return fn(s, x, y)
}

func TestAnonymousFunc() {
	result1 := test(func() int { return 100 })

	result2 := format(func(s string, x, y int) string {
		return fmt.Sprintf(s, x, y)
	}, "%d, %d", 10, 20)

	fmt.Println(result1, result2)
}

// 在Go里面，函数可以像普通变量一样被传递或使用，Go语言支持随时在代码里定义匿名函数
// 匿名函数有一个不带函数名的函数声明和函数体组成，匿名函数的优越性在于可以直接使用函数内的变量，不必声明
func AnonymousFunc() {
	getSqrt := func(a float64) float64 {
		return math.Sqrt(a)
	}
	fmt.Println(getSqrt(4))
}

func A() func() int {
	i := 0
	b := func() int {
		i++
		fmt.Println(i)
		return i
	}
	return b
}

func ClosureAdd(base int) func(int) int {
	return func(i int) int {
		base += i
		return base
	}
}

func ClosureTest(base int) (func(int) int, func(int) int) {
	add := func(i int) int {
		base += i
		return base
	}
	sub := func(i int) int {
		base -= i
		return base
	}
	return add, sub
}

// 求阶乘
func Factorial(n int) int {
	if n <= 1 {
		return 1
	}
	return n * Factorial(n-1)
}

func Fibonaci(n int) int {
	if n == 0 {
		return 0
	} else if n == 1 {
		return 1
	} else {
		return Fibonaci(n-1) + Fibonaci(n-2)
	}
}
