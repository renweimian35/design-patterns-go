// Package strategy 策略模式
package strategy

import "fmt"

type Strategy interface {
	Execute()
}

// NationalRoad 具体策略,国道
type NationalRoad struct {
}

func (n *NationalRoad) Execute() {
	fmt.Println("走国道，免费，可能堵车")
}

// Expy 具体策略,高速
type Expy struct {
}

func (e *Expy) Execute() {
	fmt.Println("走高速，收费，快，可能也堵车")
}

// Context 上下文对象，维护一个对策略对象的引用
type Context struct {
	strategy Strategy
}

func (c *Context) SetStrategy(strategy Strategy) {
	c.strategy = strategy
}

func (c *Context) ExecuteStrategy() {
	c.strategy.Execute()
}
