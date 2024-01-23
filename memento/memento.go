// Package memento 备忘录模式
package memento

import "fmt"

// Memento 备忘录
type Memento struct {
	State string // 状态
	Item  string // 记录事情
}

// Originator 发起人
type Originator struct {
	State string
	Item  string
}

func (o *Originator) CreateMemento() *Memento {
	return &Memento{State: o.State, Item: o.Item}
}

func (o *Originator) SetMemento(memento *Memento) {
	o.State = memento.State
	o.Item = memento.Item
}

func (o *Originator) ShowState() {
	fmt.Printf("Current State:%s , Item:%s", o.State, o.Item)
}

// Caretaker 管理者
type Caretaker struct {
	Memento *Memento
}
