package mediator

import (
	"fmt"
	"strings"
)

// 中介者模式封装对象之间的交互，使依赖变得简单，并且使复杂交互简单化，封装在中介者中
// 例子中的中介者使用单例模式生成中介者
// 中介者的change使用switch判断类型

var mediator *Mediator

type Mediator struct {
	CD    *CDDriver
	CPU   *CPU
	Video *VideoCard
	Sound *SoundCard
}

type CDDriver struct {
	Data string
}

type CPU struct {
	Video string
	Sound string
}

type VideoCard struct {
	Data string
}

type SoundCard struct {
	Data string
}

func (driver *CDDriver) ReadData() {
	driver.Data = "music,image"

	fmt.Printf("CDDriver: reading data %s\n", driver.Data)
	GetMediatorInstance().changed(driver)
}

func (c *CPU) Process(data string) {
	sp := strings.Split(data, ",")
	c.Sound = sp[0]
	c.Video = sp[1]

	fmt.Printf("CPU: split data with Sound %s, Video %s\n", c.Sound, c.Video)
	GetMediatorInstance().changed(c)
}

func (v *VideoCard) Display(data string) {
	v.Data = data
	fmt.Printf("VideoCard: display %s\n", v.Data)
	//GetMediatorInstance().changed(v)
}

func (s *SoundCard) Play(data string) {
	s.Data = data
	fmt.Printf("SoundCard: play %s\n", s.Data)
	//GetMediatorInstance().changed(s)
}

func GetMediatorInstance() *Mediator {
	if mediator == nil {
		mediator = &Mediator{}
	}
	return mediator
}

func (mediator *Mediator) changed(i interface{}) {
	switch inst := i.(type) {
	case *CDDriver:
		mediator.CPU.Process(inst.Data)
	case *CPU:
		mediator.Sound.Play(inst.Sound)
		mediator.Video.Display(inst.Video)
	}
}
