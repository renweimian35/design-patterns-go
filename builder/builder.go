// Package builder 建造者模式
package builder

// Car 产品类
type Car struct {
	Wheel    string // 车轮子
	Steering string // 方向盘
	Seat     string // 车座椅

}

// Builder 抽象建造者接口
type Builder interface {
	BuildWheel()
	BuildSteering()
	BuildSeat()
	GetResult() Car
}

// ConcreteCar 具体建造者类
type ConcreteCar struct {
	car Car
}

func (c *ConcreteCar) BuildWheel() {
	c.car.Wheel = "生产轮子，使用橡胶、胶水、用模具灌溉等等。"

}
func (c *ConcreteCar) BuildSteering() {
	c.car.Steering = "生产方向盘，使用模具，用塑料等产物。"

}
func (c *ConcreteCar) BuildSeat() {
	c.car.Seat = "生产座椅，用吕啊钢啊布啊等等柔和而成。"
}

func (c *ConcreteCar) GetResult() Car {
	return c.car
}

// Director 指导者类
type Director struct {
	builder Builder
}

// Construct 指导者来组合产品属性，拼装车
func (d *Director) Construct() Car {
	d.builder.BuildWheel()
	d.builder.BuildSteering()
	d.builder.BuildSeat()
	return d.builder.GetResult()
}
