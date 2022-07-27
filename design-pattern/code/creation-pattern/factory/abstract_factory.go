package factory

import "fmt"

// 抽象工厂模式：
// 提供一个创建一系列相关或者相互依赖的对象的接口，而无需指定他们的具体类
// 当产品只有一个的时候，抽象工厂就退化成工厂方法
type PCFactory interface {
	ProduceMouse() Mouse
	ProduceKeybo() Keybo
}

type Mouse interface {
	LeftClick()
	RightClick()
}

type Keybo interface {
	Enter()
	Delete()
}

type DellPcFactory struct{}

type HpPcFactory struct{}

type DellMouse struct{}

type HpMouse struct{}

type DellKeybo struct{}

type HpKeybo struct{}

func (DellPcFactory) ProduceMouse() Mouse {
	return &DellMouse{}
}

func (HpPcFactory) ProduceMouse() Mouse {
	return &HpMouse{}
}

func (DellPcFactory) ProduceKeybo() Keybo {
	return &DellKeybo{}
}

func (HpPcFactory) ProduceKeybo() Keybo {
	return &HpKeybo{}
}

func (DellMouse) LeftClick() {
	fmt.Println("dell mouse left click")
}

func (DellMouse) RightClick() {
	fmt.Println("dell mouse right click")
}

func (HpMouse) LeftClick() {
	fmt.Println("hp mouse left click")
}

func (HpMouse) RightClick() {
	fmt.Println("hp mouse right click")
}

func (DellKeybo) Enter() {
	fmt.Println("dell keyboard click enter")
}

func (DellKeybo) Delete() {
	fmt.Println("dell keyboard click delete")
}

func (HpKeybo) Enter() {
	fmt.Println("hp keyboard click enter")
}

func (HpKeybo) Delete() {
	fmt.Println("hp keyboard click delete")
}
