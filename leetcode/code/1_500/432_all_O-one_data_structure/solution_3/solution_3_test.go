package solution_3

import (
	"fmt"
	"testing"
)

func TestConstructor(t *testing.T) {
	allOne := Constructor()
	allOne.Inc("hello")
	allOne.Inc("hello")
	fmt.Println(allOne.GetMaxKey())
	fmt.Println(allOne.GetMinKey())
	//fmt.Printf("%v\n", allOne.index)
	//fmt.Printf("%v\n", allOne.maxAndMin)
	allOne.Inc("leet")
	fmt.Printf("%v\n", allOne.index)
	fmt.Println(allOne.GetMaxKey())
	fmt.Println(allOne.GetMinKey())
}

func TestConstructor1(t *testing.T) {
	allOne := Constructor()
	allOne.Inc("hello")
	allOne.Inc("goodbye")
	allOne.Inc("hello")
	allOne.Inc("hello")
	fmt.Println(allOne.GetMaxKey())
	allOne.Inc("leet")
	allOne.Inc("code")
	allOne.Inc("leet")

	allOne.Dec("hello")
	allOne.Inc("leet")
	allOne.Inc("code")
	allOne.Inc("code")
	fmt.Println(allOne.GetMaxKey())
}

func TestConstructor2(t *testing.T) {
	allOne := Constructor()
	allOne.Inc("a")
	allOne.Inc("b")
	allOne.Inc("c")
	allOne.Inc("d")
	allOne.Inc("a")
	allOne.Inc("b")
	allOne.Inc("c")
	allOne.Inc("d")
	allOne.Inc("c")
	allOne.Inc("d")
	allOne.Inc("d")
	allOne.Inc("a")
	fmt.Println(allOne.GetMinKey())
	fmt.Println(allOne.GetMaxKey())
}
