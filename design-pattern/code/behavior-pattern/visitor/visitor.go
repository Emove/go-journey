package visitor

import "fmt"

// 访问者模式可以给一系列对象透明的添加功能，并且把相关代码封装到一个类中
// 对象只要预留访问者接口Accept,则后期为对象添加功能的时候就不需要改动对象

type Customer interface {
	Accept(Visitor)
}

type Visitor interface {
	Visit(Customer)
}

type EnterpriseCustomer struct {
	name string
}

type CustomerCol struct {
	Customers []Customer
}

func (col *CustomerCol) Add(cus Customer) {
	col.Customers = append(col.Customers, cus)
}

func (col *CustomerCol) Accept(visitor Visitor) {
	for _, customer := range col.Customers {
		customer.Accept(visitor)
	}
}

func NewEnterpriseCustomer(name string) *EnterpriseCustomer {
	return &EnterpriseCustomer{
		name: name,
	}
}

func (cus *EnterpriseCustomer) Accept(visitor Visitor) {
	visitor.Visit(cus)
}

type IndividualCustomer struct {
	name string
}

func NewIndividualCustomer(name string) *IndividualCustomer {
	return &IndividualCustomer{
		name: name,
	}
}

func (cus *IndividualCustomer) Accept(visitor Visitor) {
	visitor.Visit(cus)
}

type ServiceRequestVisitor struct {
}

func (*ServiceRequestVisitor) Visit(customer Customer) {
	switch c := customer.(type) {
	case *EnterpriseCustomer:
		fmt.Printf("serving enterprise customer %s\n", c.name)
	case *IndividualCustomer:
		fmt.Printf("serving individual custoemr %s\n", c.name)
	}
}

type AnalysisVisitor struct {
}

func (*AnalysisVisitor) Visit(customer Customer) {
	switch c := customer.(type) {
	case *EnterpriseCustomer:
		fmt.Printf("analysis enterprise customer %s \n", c.name)
	}
}
