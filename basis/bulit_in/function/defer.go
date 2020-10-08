package function

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

/*
 * @author: Emove
 * @date: Do not edit
 */
func Multi(x, y int) (z int) {
	defer func() {
		z *= 10
	}()
	z = x + y
	return z
}

type User struct {
	name string
}

func (user *User) Close() {
	fmt.Println(user.name, "closed")
}

func ErrorTest() {
	users := []User{
		{"a"}, {"b"}, {"c"},
	}
	for _, user := range users {
		defer user.Close()
	}
}

func CorrectTest() {
	users := []User{
		{"a"}, {"b"}, {"c"},
	}
	for _, user := range users {
		u := user
		defer u.Close()
	}
}

func ErrorOccurWherExecuteDeferTest(x int) {
	defer fmt.Println("a")
	defer fmt.Println("b")
	defer func() {
		fmt.Println(100 / x)
	}()
	defer fmt.Println("c")
}

func DeferTest() {
	x, y := 10, 20
	defer func(i int) {
		fmt.Println("defer:", i, y)
	}(x) // x被复制
	x += 10
	y += 100
	fmt.Println("x = ", x, "y = ", y)
}

func runTest() {
	var run func() = nil
	defer run()
	fmt.Println("runs...")
}

func ErrTest() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
		}
	}()
	runTest()
}

func CloseHttpBodyTest() {
	res, err := http.Get("http://www.baidu.com")
	if res != nil {
		// 当且仅当get成功时才执行defer
		defer res.Body.Close()
	}
	if err != nil {
		fmt.Println(err)
	}
}

func CloseFileTest() {
	f, err := os.Open("book.txt")
	if err != nil {
		fmt.Println(err)
	}
	if f != nil {
		defer func(f io.Closer) {
			if err := f.Close(); err != nil {
				fmt.Println("defer close book.txt err %v\n", err)
			}
		}(f)
	}

	f, err = os.Open("anthor-boox.txt")
	if err != nil {
		fmt.Println(err)
	}
	if f != nil {
		defer func(f io.Closer) {
			if err := f.Close(); err != nil {
				fmt.Println("defer close anthor-book.txt err %v\n", err)
			}
		}(f)
	}
}
