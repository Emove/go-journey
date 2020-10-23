package factory

import (
	"fmt"
	"testing"
)

func TestSingletonBean_GetObject(t *testing.T) {
	beanFactory := &SingletonBeanFactory{}
	singletonBean := beanFactory.GetBean()
	object1 := singletonBean.GetObject()
	object2 := singletonBean.GetObject()
	if object1 != object2 {
		t.Fatal("object1 is not equals object2")
	}
}

func TestPrototypeBean_GetObject(t *testing.T) {
	beanFactory := &PrototypeBeanFactory{}
	bean := beanFactory.GetBean()
	object1 := bean.GetObject()
	object2 := bean.GetObject()
	fmt.Printf("object1 : %#p \n", object1)
	fmt.Printf("object2 : %#p \n", object2)
	if object1 == object2 {
		t.Fatal("object1 is equals object2")
	}
}
