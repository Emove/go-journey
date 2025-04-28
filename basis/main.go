/*
 * @author: Emove
 * @date: Do not edit
 */
package main

import (
	"basis/bulit_in/code"
	"basis/bulit_in/function"
	"basis/bulit_in/method"
	"basis/bulit_in/oop"
	"basis/bulit_in/time"
	"fmt"
)

func main() {
	fmt.Println("this is main function")
	// testConsts()
	// testIota()
	//testSlice()
	//testMap()
	// testStruct()
	// testSwitch()
	// testFunction()
	// testDefer()
	// testError()
	//testMethod()
	//testOo()
	//testSelect()
	//TestChannel()
	//TestTime()
	//print(generic.IsValidValue("Bool", "sda"))
	//fmt.Println(1 == nil)
}

func TestTime() {
	time.GetYesterdayTime()
}

func testConsts() {
	code.ArrayPrint()
}

func testIota() {
	code.PrintDemo1()
	code.PrintSkip1()
	code.PrintSkip2()
	code.PrintByteUnits()
	code.PrintMultiAssignment()
}

func testSlice() {
	//code.CreateSlice()
	// code.PrintVariable()
	// code.PrintCreateByMake()
	// code.PrintModify()
	// code.PrintInit()
	// code.AppendOutOfCapTest()
	// code.SliceAssignRuleTest()
	// code.SliceCopyTest()
	// code.ForEachSliceTest()
	// code.SliceResizeTest()
	// code.SliceStringTest()
	//code.SliceChangeStringTest2()
	//code.AppendAndForeachTest()
	//code.SplitSlice()
	code.OutOfCap()
}

func testMap() {
	// code.IfExistsTest()
	// code.ForEachMapTest()
	code.DeleteFormMapTest()
	// code.MapSliceTest()
	//code.SliceMapTest()
}

func testStruct() {
	// code.PrintPersonInstance()
	// code.CreatePersonInstanceByNew()
	// code.CreatePersonInstance()
	// code.PutPersonInfo2Map()
	// per := code.Construct("Emove", "sz", 20)
	// per.Dream()
	// per.SetAge(18)
	// fmt.Printf("person = %#v\n", per)
	// code.NewNestedStruct()
	dog := code.NewDogStruct()
	dog.Yelp()
	dog.Move()
	fmt.Printf("dog = %#v\n", dog)
}

func testSwitch() {
	code.Test()
}

func testFunction() {
	// function.TestAnonymousFunc()
	// result := function.Multi(1, 2)
	// fmt.Println(result)
	// function.AnonymousFunc()
	// c := function.A()
	// c()
	// c()
	// c()
	// function.A()
	// c()
	// temp1 := function.ClosureAdd(10)
	// fmt.Println(temp1(1), temp1(2))
	// add, sub := function.ClosureTest(10)
	// fmt.Println(add(1), sub(2))
	// fmt.Println(add(2), sub(2))
	// result := function.Factorial(7)
	// fmt.Println(result)
	result := function.Fibonaci(7)
	fmt.Println(result)
}

func testDefer() {
	// function.ErrorTest()
	// function.CorrectTest()
	// function.ErrorOccurWherExecuteDeferTest(0)
	// function.DeferTest()
	// function.ErrTest()
	// function.CloseHttpBodyTest()
	function.CloseFileTest()
}

func testError() {
	// function.PanicTest()
	// function.PanicWhenSendData2ClosedChannel()
	// function.MultiDeferPanicTest()
	// function.MultiDeferPanicTest2()
	// function.AnonymousDeferTest()
	// function.AnonymousFuncTest(2, 2)
	// defer func() {
	// 	fmt.Println(recover())
	// }()
	// switch z, err := function.ErrorWhenDivTest(10, 1); err {
	// case nil:
	// 	fmt.Println(z)
	// case function.ErrorDivByZero:
	// 	panic(err)
	// }
	// function.Try(func() {
	// 	fmt.Println("执行try代码块")
	// 	fmt.Println("抛出异常panic")
	// 	panic("throw panic")
	// }, func(err interface{}) {
	// 	fmt.Printf("catch error:%#v\n", err)
	// })
	// area, err := function.GetCircleArea(-2)
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// fmt.Println(area)

	err := function.TryOpen("book.txt")
	if err != nil {
		switch v := err.(type) {
		case *function.PathError:
			fmt.Println(v)
		default:
		}
	}
}

func testMethod() {
	// user := &method.User{
	// 	Name: "emove",
	// 	Age:  20,
	// }
	// fmt.Printf("address: %p\n", user)
	// user.Values()
	// user.Pointer()
	manager := method.Manager{
		method.User{
			Name: "Emove",
			Age:  20,
		},
	}
	manager.Values()
	manager.Pointer()
	m1 := &manager
	m1.Values()
	m1.Pointer()

	/*// 隐式传递reciver
	valuesMethod := m1.Values
	valuesMethod()

	// 显式传递reciver
	pointerMehtod := (*method.User).Pointer
	pointerMehtod(&method.User{
		Name: "Emove",
		Age:  20,
	})*/
}

func testOo() {
	oop.NewStudentTest()
}

func testSelect() {
	ch := make(chan int)
	go code.BreakSelectTest(ch)

	for i := 0; i < 10; i++ {
		ch <- i
	}
}

func TestChannel() {
	code.LenOfChannel()
}
