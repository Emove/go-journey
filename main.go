package main

import (
	"fmt"
	"go-demo/basis/bulit_in/code"
)

func main() {
	fmt.Println("this is main function")
	// testConsts()
	// testIota()
	// testSlice()
	// testMap()
	testStruct()
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
	code.PrintPersonInstance()
	code.CreatePersonInstanceByNew()
}