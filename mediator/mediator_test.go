package mediator

import "testing"

func TestMediator(t *testing.T) {
	// 创建聊天室中介者
	chatMediator := NewConcreteMediator()

	// 创建用户并注册到中介者
	user1 := NewConcreteUser("User1", chatMediator)
	user2 := NewConcreteUser("User2", chatMediator)
	user3 := NewConcreteUser("User3", chatMediator)

	chatMediator.RegisterUser(user1)
	chatMediator.RegisterUser(user2)
	chatMediator.RegisterUser(user3)

	// 用户发送消息
	user1.SendMessage("Hello, 我想问下，我有两百块钱可以买什么车？")
	user2.SendMessage("Hi, User1!你可以买个锤子")
}
