package prototype

type Cloneable interface {
	Clone() Cloneable
}

type PrototypeManager struct {
	prototypes map[string]Cloneable
}

func NewPrototypeManager() *PrototypeManager {
	return &PrototypeManager{
		prototypes: make(map[string]Cloneable),
	}
}

func (manager *PrototypeManager) Get(name string) Cloneable {
	return manager.prototypes[name]
}

func (manager *PrototypeManager) Set(name string, prototype Cloneable) {
	manager.prototypes[name] = prototype
}
