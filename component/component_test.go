package component

import "testing"

func TestComponent(t *testing.T) {
	// 创建车辆组合
	vehicle := NewComposite("Car")

	// 添加部件
	engine := &Leaf{Name: "Engine"}
	tire := &Leaf{Name: "Tire"}
	seat := &Leaf{Name: "Seat"}

	vehicle.Add(engine)
	vehicle.Add(tire)
	vehicle.Add(seat)

	// 执行操作
	vehicle.Operation()
}
