// Package responsibility_chain 责任链模式
package responsibility_chain

import "fmt"

// Handler 处理者接口
type Handler interface {
	HandleRequest(amount int)
	SetSuccessor(successor Handler)
}

// CarHandler 购车处理者，具体处理者
type CarHandler struct {
	Price     int
	Car       string
	successor Handler
}

func (c *CarHandler) HandleRequest(price int) {
	if price >= c.Price {
		fmt.Printf("你的金额为:%d,你可以买 %s\n", price, c.Car)
	} else if c.successor != nil {
		fmt.Printf("你的金额为:%d, 你不能买 %s\n ", price, c.Car)
		fmt.Printf("不过我们可以看看下一个你可以买的车.\n")
		c.successor.HandleRequest(price)
	} else {
		fmt.Printf("Sorry, 你的金额%d，买不了任何车，这边建议你买电动车\n", price)
	}
}

func (c *CarHandler) SetSuccessor(successor Handler) {
	c.successor = successor
}
