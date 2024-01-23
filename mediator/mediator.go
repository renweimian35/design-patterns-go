// Package mediator 中介模式
package mediator

import "fmt"

// Mediator 中介者接口
type Mediator interface {
	SendMessage(message string, user User)
}

// ConcreteMediator 具体中介者
type ConcreteMediator struct {
	users map[string]User
}

func NewConcreteMediator() *ConcreteMediator {
	return &ConcreteMediator{
		users: make(map[string]User),
	}
}

func (c *ConcreteMediator) RegisterUser(user User) {
	c.users[user.GetName()] = user
}

func (c *ConcreteMediator) SendMessage(message string, user User) {
	fmt.Printf("[%s] sends message: %s\n", user.GetName(), message)
	for _, u := range c.users {
		if u != user {
			u.ReceiveMessage(message)
		}
	}
}

// User 用户接口
type User interface {
	SendMessage(message string)
	ReceiveMessage(message string)
	GetName() string
}

// ConcreteUser 具体用户
type ConcreteUser struct {
	name     string
	mediator Mediator
}

func NewConcreteUser(name string, mediator Mediator) *ConcreteUser {
	return &ConcreteUser{
		name:     name,
		mediator: mediator,
	}
}

func (u *ConcreteUser) SendMessage(message string) {
	u.mediator.SendMessage(message, u)
}

func (u *ConcreteUser) ReceiveMessage(message string) {
	fmt.Printf("[%s] receives message: %s\n", u.name, message)
}

func (u *ConcreteUser) GetName() string {
	return u.name
}
