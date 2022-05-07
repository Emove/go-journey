package builder

import (
	"fmt"
)

type Builder interface {
	PartA()
	PartB()
	PartC()
}

type Director struct {
	builder Builder
}

func NewDirector(builder Builder) *Director {
	return &Director{
		builder: builder,
	}
}

func (director *Director) Constructor() {
	director.builder.PartA()
	director.builder.PartB()
	director.builder.PartC()
}

type ComputerBuilder struct {
	finished bool
}

func (builder *ComputerBuilder) PartA() {
	if !builder.finished {
		fmt.Println("the main board assemble finish")
	}
}

func (builder *ComputerBuilder) PartB() {
	if !builder.finished {
		fmt.Println("the battery assemble finish")
	}
}

func (builder *ComputerBuilder) PartC() {
	fmt.Println("the crate assemble finish")
	builder.finished = true
}
