/*
 * @author: Emove
 * @date: Do not edit
 */
/*
 * @author: Emove
 * @date: Do not edit
 */
package main

import (
	"fmt"
	"go-demo/basis/bulit_in/code"
	"go-demo/basis/bulit_in/function"
)

func main() {
	fmt.Println("this is main function")
	// testConsts()
	// testIota()
	// testSlice()
	// testMap()
	// testStruct()
	// testSwitch()
	// testMethod()
	testDefer()
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
	// code.CreateSlice()
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
	code.SliceChangeStringTest2()
}

func testMap() {
	// code.IfExistsTest()
	// code.ForEachMapTest()
	// code.DeleteFormMapTest()
	// code.MapSliceTest()
	code.SliceMapTest()
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

func testMethod() {
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
