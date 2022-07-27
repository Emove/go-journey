package factory

import (
	"design-pattern/code/creation-pattern/singleton"
)

type Bean interface {
	GetObject() interface{}
}

type BeanFactory interface {
	GetBean() Bean
}

type SingletonBeanFactory struct {
}

type PrototypeBeanFactory struct {
}

type SingletonBean struct {
}

type PrototypeBean struct {
}

func (singletonBean SingletonBean) GetObject() interface{} {
	return singleton.GetInstance()
}

func (prototypeBean PrototypeBean) GetObject() interface{} {
	return &Instance{
		Id: "232131232",
	}
}

func (beanFactory SingletonBeanFactory) GetBean() Bean {
	return &SingletonBean{}
}

func (BeanFactory PrototypeBeanFactory) GetBean() Bean {
	return &PrototypeBean{}
}

type Instance struct {
	Id string
}
