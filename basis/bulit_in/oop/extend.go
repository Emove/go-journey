/**
 * @author Emove
 * @date 2021/1/19
 */
package oop

import "fmt"

type Function interface {
	Apply()
	DoApply()
}

type BaseFunction struct {
}

func (base *BaseFunction) Apply() {
	fmt.Println("base function apply called")
	base.DoApply()
}

func (base *BaseFunction) DoApply() {
	fmt.Println("base function apply called")
}

type ConcreteFunction struct {
	BaseFunction
}

func (concrete *ConcreteFunction) Apply() {
	concrete.BaseFunction.Apply()
}

func (concrete *ConcreteFunction) DoApply() {
	fmt.Println("concrete function doApply called")
}
