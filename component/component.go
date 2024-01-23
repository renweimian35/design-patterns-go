package component

import "fmt"

// Component 组件接口
type Component interface {
	Operation()
}

// Leaf 叶子节点（车辆部件）
type Leaf struct {
	Name string
}

func (l *Leaf) Operation() {
	fmt.Printf("Part: %s\n", l.Name)
}

// Composite 容器节点（车辆）
type Composite struct {
	Name     string
	Children []Component
}

func NewComposite(name string) *Composite {
	return &Composite{Name: name, Children: make([]Component, 0)}
}

func (c *Composite) Add(component Component) {
	c.Children = append(c.Children, component)
}

func (c *Composite) Operation() {
	fmt.Printf("Vehicle: %s\n", c.Name)
	for _, child := range c.Children {
		child.Operation()
	}
}
