package adapter

// Target 适配的目标接口
type Target interface {
	Request() string
}

// Adaptee 被适配的接口
type Adaptee interface {
	SpecificRequest() string
}

type AdapteeImpl struct {
}

type Adapter struct {
	adaptee Adaptee
}

func (*AdapteeImpl) SpecificRequest() string {
	return "adaptee method"
}

func (adapter *Adapter) Request() string {
	return adapter.adaptee.SpecificRequest()
}

// NewAdaptee 配适配接口的工厂函数
func NewAdaptee() Adaptee {
	return &AdapteeImpl{}
}

func NewAdapter(adaptee Adaptee) Target {
	return &Adapter{adaptee: adaptee}
}
