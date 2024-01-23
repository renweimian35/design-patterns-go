// Package bridge 桥接模式
package bridge

import "fmt"

// ColorImplementor 颜色的实现接口
type ColorImplementor interface {
	Colorize() string
}

// ColorRed 具体的颜色实现类：红色
type ColorRed struct {
}

func (c ColorRed) Colorize() string {
	return "Red"
}

// ColorBlue 具体的颜色实现类：蓝色
type ColorBlue struct {
}

func (c ColorBlue) Colorize() string {
	return "Blue"
}

// Shape 形状的抽象类
type Shape interface {
	Hub() string                                      // 轮毂型号
	SetColorImplementor(implementor ColorImplementor) // 喷漆颜色
}

// Large 大号的轮毂
type Large struct {
	colorImplementor ColorImplementor
}

func (l *Large) Hub() string {
	return fmt.Sprintf("hub size large:8.0J,color:%s", l.colorImplementor.Colorize())
}
func (l *Large) SetColorImplementor(implementor ColorImplementor) {
	l.colorImplementor = implementor
}

// Medium 中号的轮毂
type Medium struct {
	colorImplementor ColorImplementor
}

func (m *Medium) Hub() string {
	return fmt.Sprintf("hub size medium:7.0J,color:%s", m.colorImplementor.Colorize())
}
func (m *Medium) SetColorImplementor(implementor ColorImplementor) {
	m.colorImplementor = implementor
}
