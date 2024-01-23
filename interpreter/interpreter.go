package interpreter

import "strings"

// Car 车辆对象
type Car struct {
	Model string
	Color string
}

// Context 上下文，包含解释器需要的全局信息
type Context struct {
	car *Car
}

// Expression 表达式接口
type Expression interface {
	Interpret(context *Context)
}

// ModelExpression 车型表达式
type ModelExpression struct {
	model string
}

func (m *ModelExpression) Interpret(context *Context) {
	context.car.Model = m.model
}

// ColorExpression 颜色表达式
type ColorExpression struct {
	color string
}

func (c *ColorExpression) Interpret(context *Context) {
	context.car.Color = c.color
}

// Parser 解析器，根据输入字符串构建车辆配置
type Parser struct {
	expressions []Expression
}

func (p *Parser) Parse(description string) {
	config := strings.Split(description, " ")

	for _, item := range config {
		switch {
		case strings.Contains(item, "Bus"):
			p.expressions = append(p.expressions, &ModelExpression{"Bus"})
		case strings.Contains(item, "Red"):
			p.expressions = append(p.expressions, &ColorExpression{"Red"})
		}
	}
}

func (p *Parser) Interpret(context *Context) {
	for _, expression := range p.expressions {
		expression.Interpret(context)
	}
}
