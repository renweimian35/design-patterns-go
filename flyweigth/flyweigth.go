// Package flyweigth 享元模式
package flyweigth

import "fmt"

// Flyweight 享元接口
type Flyweight interface {
	Operation(extrinsicState string)
}

// ConcreteFlyweight 具体享元
type ConcreteFlyweight struct {
	intrinsicState string
}

func NewConcreteFlyweight(intrinsicState string) *ConcreteFlyweight {
	return &ConcreteFlyweight{intrinsicState: intrinsicState}
}

func (c *ConcreteFlyweight) Operation(extrinsicState string) {
	fmt.Printf("Operation with intrinsic state %s and extrinsic state %s\n", c.intrinsicState, extrinsicState)
}

// FlyweightFactory 享元工厂
type FlyweightFactory struct {
	flyweights map[string]Flyweight
}

func NewFlyweightFactory() *FlyweightFactory {
	return &FlyweightFactory{
		flyweights: make(map[string]Flyweight),
	}
}

func (f *FlyweightFactory) GetFlyweight(key string) Flyweight {
	if flyweight, ok := f.flyweights[key]; ok {
		return flyweight
	} else {
		newFlyweight := NewConcreteFlyweight(key)
		f.flyweights[key] = newFlyweight
		return newFlyweight
	}
}
