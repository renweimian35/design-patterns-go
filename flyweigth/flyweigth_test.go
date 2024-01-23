package flyweigth

import (
	"fmt"
	"testing"
)

func TestFlyweigth(t *testing.T) {
	// 创建享元工厂
	factory := NewFlyweightFactory()

	// 获取或创建享元对象
	flyweight1 := factory.GetFlyweight("shared")
	flyweight2 := factory.GetFlyweight("shared")

	// 客户端操作
	flyweight1.Operation("state1")
	flyweight2.Operation("state2")

	// 检查是否共享同一个对象
	fmt.Println(flyweight1 == flyweight2) // 输出：true
}
