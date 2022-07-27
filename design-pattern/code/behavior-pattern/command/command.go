package command

import "fmt"

// 命令模式本质上是把某个对象的方法调用封装到对象中，方便传递，存储和调用
// 实例中把主板单中的启动（start）方法和重启（reboot）方法封装为命令对象，再传递到主机（box）对象中，于两个按钮进行绑定：
// 第一个机箱（box1）设置按钮1（button1）为开机键，按钮2（button2）为重启
// 第二个机箱（box2）设置按钮2（button2）为开机键，按钮1（button1）为重启
// 从而得到配置的灵活性，除了配置灵活性外，使用命令模式还可以用作批处理，任务队列，undo、redo等

type Command interface {
	Execute()
}

type StartCommand struct {
	mb *MotherBoard
}

func NewStartCommand(mb *MotherBoard) *StartCommand {
	return &StartCommand{
		mb: mb,
	}
}

func (c *StartCommand) Execute() {
	c.mb.Start()
}

type RebootCommand struct {
	mb *MotherBoard
}

func NewRebootCommand(mb *MotherBoard) *RebootCommand {
	return &RebootCommand{
		mb: mb,
	}
}

func (c *RebootCommand) Execute() {
	c.mb.Reboot()
}

type MotherBoard struct {
}

func (*MotherBoard) Start() {
	fmt.Println("system starting")
}

func (*MotherBoard) Reboot() {
	fmt.Println(" system rebooting")
}

type Box struct {
	button1 Command
	button2 Command
}

func NewBox(button1, button2 Command) *Box {
	return &Box{
		button1: button1,
		button2: button2,
	}
}

func (b *Box) PressButton1() {
	b.button1.Execute()
}

func (b *Box) PressButton2() {
	b.button2.Execute()
}
