package chain

import "fmt"

// 职责链模式用于分离不同职责，并且动态组合相关职责
// Golang实现职责链模式的时候，因为没有继承的支持，使用链对象包含职责的方式
// 即：
// 1、链对象包含当前职责对象以及下一个职责链
// 2、职责对象提供接口表示是否能处理对应的请求
// 3、职责对象提供处理函数处理相关职责
// 同时可在职责链类中实现职责接口相关函数，是职责对象可以当作一般职责对象使用

type Manager interface {
	HaveRight(money int) bool
	HandleFeeRequest(name string, money int) bool
}

type RequestChain struct {
	Manager
	successor *RequestChain
}

func (request *RequestChain) SetSuccessor(chain *RequestChain) {
	request.successor = chain
}

func (request *RequestChain) HandleFeeRequest(name string, money int) bool {
	if request.Manager.HaveRight(money) {
		return request.Manager.HandleFeeRequest(name, money)
	}
	if request.successor != nil {
		return request.successor.HandleFeeRequest(name, money)
	}
	return false
}

func (request *RequestChain) HaveRight(money int) bool {
	return true
}

type ProjectManager struct {
}

func NewProjectManagerChain() *RequestChain {
	return &RequestChain{
		Manager: &ProjectManager{},
	}
}

func (*ProjectManager) HaveRight(money int) bool {
	return money < 500
}

func (*ProjectManager) HandleFeeRequest(name string, money int) bool {
	if name == "bob" {
		fmt.Printf("project manager permit %s %d fee request\n", name, money)
		return true
	}
	fmt.Printf("project manager dont permit %s %d fee request\n", name, money)
	return false
}

type DepManager struct {
}

func NewDepManagerChain() *RequestChain {
	return &RequestChain{
		Manager: &DepManager{},
	}
}

func (*DepManager) HaveRight(money int) bool {
	return money < 5000
}

func (*DepManager) HandleFeeRequest(name string, money int) bool {
	if name == "tom" {
		fmt.Printf("Dep manager permit %s %d fee request\n", name, money)
		return true
	}
	fmt.Printf("Dep manager dont permit %s %d fee request\n", name, money)
	return false
}

type GeneralManager struct{}

func NewGeneralManagerChain() *RequestChain {
	return &RequestChain{
		Manager: &GeneralManager{},
	}
}

func (*GeneralManager) HaveRight(money int) bool {
	return true
}

func (*GeneralManager) HandleFeeRequest(name string, money int) bool {
	if name == "ada" {
		fmt.Printf("General manager permit %s %d fee request\n", name, money)
		return true
	}
	fmt.Printf("General manager don't permit %s %d fee request\n", name, money)
	return false
}
