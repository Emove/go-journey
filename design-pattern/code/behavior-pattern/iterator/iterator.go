package iterator

import "fmt"

// 迭代器模式用于使用相同方式迭代不同类型集合或者隐藏集合类型的具体实现
// 可以使用迭代器模式是遍历同时应用迭代策略，如请求新对象，过滤，处理对象等

type Iterator interface {
	First()
	IsDone() bool
	Next() interface{}
}

type Aggregate interface {
	Iterator() Iterator
}

type Numbers struct {
	start, end int
}

type NumbersIterator struct {
	numbers *Numbers
	next    int
}

func NewNumbers(start, end int) *Numbers {
	return &Numbers{
		start: start,
		end:   end,
	}
}

func (num *Numbers) Iterator() Iterator {
	return &NumbersIterator{
		numbers: num,
		next:    num.start,
	}
}

func (ni *NumbersIterator) First() {
	ni.next = ni.numbers.start
}

func (ni *NumbersIterator) IsDone() bool {
	return ni.next > ni.numbers.end
}

func (ni *NumbersIterator) Next() interface{} {
	if !ni.IsDone() {
		next := ni.next
		ni.next++
		return next
	}
	return nil
}

func IteratorPrint(i Iterator) {
	for i.First(); !i.IsDone(); {
		c := i.Next()
		fmt.Println(c)
	}
}
