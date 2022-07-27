package function

import (
	"errors"
	"fmt"
	"os"
	"time"
)

func PanicTest() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err.(string))
		}
	}()
	panic("panic error")
}

// 向已关闭的通道发送数据会引发panic
func PanicWhenSendData2ClosedChannel() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Printf("error : %#v\n", err)
		}
	}()
	var ch chan int = make(chan int, 10)
	close(ch)
	ch <- 1
}

// 延迟调用中引发的错误，可被后续延迟调用捕获
// 但仅最后一个错误可被捕获
func MultiDeferPanicTest() {
	defer func() {
		fmt.Println(recover())
	}()

	defer func() {
		panic("defer panic")
	}()

	panic("test panic")
}

func MultiDeferPanicTest2() {
	// 有效
	defer func() {
		fmt.Println(recover())
	}()
	// 无效
	defer recover()
	// 无效
	defer fmt.Println(recover())

	defer func() {
		func() {
			fmt.Println("defer inner")
			// 无效
			recover()
		}()
	}()
	panic("test panic")
}

func except() {
	fmt.Println(recover())
}

func AnonymousDeferTest() {
	// 有效
	defer except()
	panic("test panic")
}

func AnonymousFuncTest(x, y int) {
	var z int
	// 将代码块重构成匿名函数，可确保后续代码被执行
	func() {
		defer func() {
			if recover() != nil {
				z = 2
			}
		}()
		panic("test panic")
		z = x / y
		return
	}()

	fmt.Printf("x / y = %d \n", z)
}

var ErrorDivByZero = errors.New("division by zero")

func ErrorWhenDivTest(x, y int) (int, error) {
	if y == 0 {
		return 0, ErrorDivByZero
	}
	return x / y, nil
}

// 模拟Java语法中的try-catch
func Try(fun func(), handler func(interface{})) {
	defer func() {
		if err := recover(); err != nil {
			handler(err)
		}
	}()
	fun()
}

func GetCircleArea(radius float32) (area float32, err error) {
	if radius < 0 {
		err = errors.New("radius can not smaller than zero")
		return
	}
	area = 3.14 * radius * radius
	return
}

type PathError struct {
	path       string
	op         string
	createTime string
	message    string
}

func (err *PathError) Error() string {
	return fmt.Sprintf("path=%s \n op=%s \n createTime=%s \n message=%s", err.path, err.op, err.createTime, err.message)
}

func TryOpen(path string) error {
	file, err := os.Open(path)
	if err != nil {
		return &PathError{
			path:       path,
			op:         "read",
			message:    err.Error(),
			createTime: fmt.Sprintf("%v", time.Now()),
		}
	}
	defer file.Close()
	return nil
}
