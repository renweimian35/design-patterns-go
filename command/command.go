// Package command 命令模式
package command

import "fmt"

// Car Receiver 接收者接口
type Car interface {
	Lock()   // 上锁
	Unlock() // 开锁
}

// Command 命令接口
type Command interface {
	Execute()
}

// LockCommand 具体命令类：上锁命令
type LockCommand struct {
	Car Car
}

func (c *LockCommand) Execute() {
	c.Car.Lock()
}

type UnLockCommand struct {
	Car Car
}

func (c *UnLockCommand) Execute() {
	c.Car.Unlock()
}

// CarSystem ConcreteReceiver 具体接收者类：汽车
type CarSystem struct{}

func (c *CarSystem) Lock() {
	fmt.Println("车门上锁")
}

func (c *CarSystem) Unlock() {
	fmt.Println("车门解锁")
}

// CarOwner 调用者类：车主
type CarOwner struct {
	Command Command
}

func (o *CarOwner) PressButton() {
	o.Command.Execute()
}
