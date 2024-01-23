package decorator

import "fmt"

// Car ：Component 车接口
type Car interface {
	Drive() string // 车基本的结构、功能就是drive
}

// BasicCar ：ConcreteComponent 具体车类
type BasicCar struct{}

func (c *BasicCar) Drive() string {
	return "BasicCar: Driving"
}

// CarDecorator ：Decorator 装饰器抽象类
type CarDecorator interface {
	Car
}

// NavigationDecorator 导航系统装饰器
type NavigationDecorator struct {
	car Car // 导航装在车里
}

func (nd *NavigationDecorator) Drive() string {
	return fmt.Sprintf("NavigationDecorator: %s with Navigation System", nd.car.Drive())
}
