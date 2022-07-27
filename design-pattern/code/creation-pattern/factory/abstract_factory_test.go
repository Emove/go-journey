package factory

import "testing"

func TestDellPcFactory(t *testing.T) {
	factory := &DellPcFactory{}
	keybo := factory.ProduceKeybo()
	keybo.Enter()
	keybo.Delete()
	mouse := factory.ProduceMouse()
	mouse.LeftClick()
	mouse.RightClick()
}

func TestHpPcFactory(t *testing.T) {
	factory := &HpPcFactory{}
	mouse := factory.ProduceMouse()
	mouse.RightClick()
	mouse.LeftClick()
	keybo := factory.ProduceKeybo()
	keybo.Delete()
	keybo.Enter()
}
