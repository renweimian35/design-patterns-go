// Package state 状态模式
package state

import "fmt"

// State 抽象状态接口
type State interface {
	Handle()
}

// Car 车辆对象，维护一个对具体状态对象的引用
type Car struct {
	state State
}

func (c *Car) SetState(state State) {
	c.state = state
}

func (c *Car) Request() {
	c.state.Handle()
}

// ChildLock 具体状态A
type ChildLock struct{}

func (s *ChildLock) Handle() {
	fmt.Println("儿童锁上锁，此时车门不能打开")
}

type ChildUnLock struct {
}

func (s *ChildUnLock) Handle() {
	fmt.Println("儿童锁解锁，此时车门可以打开")
}
