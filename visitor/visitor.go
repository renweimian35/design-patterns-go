// Package visitor visitor 访问者模式
package visitor

import "fmt"

// Visitor 访问者接口
type Visitor interface {
	Visit()
}

// BMWVisitor 具体访问者
type BMWVisitor struct {
}

func (b BMWVisitor) Visit() {
	fmt.Println("我是宝马销售,我们的车是豪华品牌")
}

type MazdaVisitor struct {
}

func (b MazdaVisitor) Visit() {
	fmt.Println("我是马自达销售,我们的车，现在不堵了")
}

// IElement 元素接口
type IElement interface {
	Accept(visitor Visitor)
}

// Element 具体元素
type Element struct {
}

func (e *Element) Accept(v Visitor) {
	v.Visit()
}
