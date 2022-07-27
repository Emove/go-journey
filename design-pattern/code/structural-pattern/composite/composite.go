package composite

import "fmt"

const (
	LeafNode = iota
	CompositeNode
)

type Component interface {
	Parent() Component
	SetParent(component Component)
	Name() string
	SetName(name string)
	AddChild(component Component)
	Print(string2 string)
}

func NewComponent(kind int, name string) Component {
	var com Component
	switch kind {
	case LeafNode:
		com = NewLeaf()
	case CompositeNode:
		com = NewComposite()
	}
	com.SetName(name)
	return com
}

type component struct {
	parent Component
	name   string
}

func (c *component) Parent() Component {
	return c.parent
}

func (c *component) SetParent(com Component) {
	c.parent = com
}

func (c *component) Name() string {
	return c.name
}

func (c *component) SetName(name string) {
	c.name = name
}

func (c *component) AddChild(com Component) {

}

func (c *component) Print(string) {

}

type Leaf struct {
	component
}

func NewLeaf() *Leaf {
	return &Leaf{}
}

func (leaf *Leaf) Print(pre string) {
	fmt.Printf("%s - %s\n", pre, leaf.name)
}

type Composite struct {
	component
	children []Component
}

func NewComposite() *Composite {
	return &Composite{
		children: make([]Component, 0),
	}
}

func (c *Composite) AddChild(child Component) {
	child.SetParent(c)
	c.children = append(c.children, child)
}

func (c *Composite) Print(pre string) {
	fmt.Printf("%s + %s \n", pre, c.Name())
	pre += "  "
	for _, comp := range c.children {
		comp.Print(pre)
	}
}
