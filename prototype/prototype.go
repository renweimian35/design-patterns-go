// Package prototype 原型模式
package prototype

type Prototype interface {
	Name() string
	Clone() Prototype
}

func NewConcretePrototype(name string) *ConcretePrototype {
	return &ConcretePrototype{name: name}
}

type ConcretePrototype struct {
	name string
}

func (c *ConcretePrototype) Name() string {
	return c.name
}

func (c *ConcretePrototype) Clone() Prototype {
	return &ConcretePrototype{
		name: c.name,
	}
}
