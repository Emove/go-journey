package code

import (
	"sync"
	"testing"
)

func TestGetLastElement(t *testing.T) {
	GetLastElement()
}

func TestSliceStringTest(t *testing.T) {
	SliceStringTest()
}

func TestIfCopyWhenForeachSlice(t *testing.T) {
	IfCopyWhenForeachSlice()
}

func TestSplitSlice(t *testing.T) {
	SplitSlice()
}

func TestRemoveOnForeach(t *testing.T) {
	RemoveOnForeach()
}

func TestRemoveTheLastOne(t *testing.T) {
	RemoveTheLastOne()
}

func TestLenOfSkipSlice(t *testing.T) {
	LenOfSkipSlice()
}

func TestLenOfSlice(t *testing.T) {
	LenOfSlice()
}

func TestMoveBackwardTest(t *testing.T) {
	MoveBackwardTest()
}

func TestSliceCopyTest(t *testing.T) {
	SliceCopyTest()
}

func TestDelete(t *testing.T) {
	Delete()
}

func TestAppendToInterface(t *testing.T) {
	AppendToInterface()
}

func TestValueCopy(t *testing.T) {
	ValueCopy()
}

func TestClear(t *testing.T) {
	Clear()
}

func TestAppend2Nil(t *testing.T) {
	Append2Nil()
}

func TestAppendEmptySlice(t *testing.T) {
	AppendEmptySlice()
}

func TestAppendToNilSlice(t *testing.T) {
	AppendToNilSlice()
}

func TestRemoveTheFirstOne(t *testing.T) {
	RemoveTheFirstOne()
}

func TestPrintVariable(t *testing.T) {
	PrintVariable()
}

func TestIsRef(t *testing.T) {
	wait := &sync.WaitGroup{}
	wait1 := &sync.WaitGroup{}
	wait.Add(1)
	wait1.Add(1)
	_, s2 := IsRef(wait, wait1)
	//s1[0] = 100
	s2[0] = 100
	wait.Done()
	wait1.Wait()
}

func TestForeachNilSlice(t *testing.T) {
	ForeachNilSlice()
}
