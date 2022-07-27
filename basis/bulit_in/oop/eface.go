package oop

type Resource interface {
	TryLock() (bool, error)
	UnLock() error
}

type Node struct {
}

func (n *Node) TryLock() (bool, error) {
	return false, nil
}

func (n *Node) UnLock() error {
	return nil
}

func NewNilNode() *Node {
	return nil
}
