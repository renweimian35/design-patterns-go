// Package iterator 迭代器模式
package iterator

type Iterator interface {
	HasNext() bool
	Next() interface{}

	HasPrev() bool
	Prev() interface{}
}

type Aggregate interface {
	CreateIterator() Iterator
}

// ConcreteAggregate 具体聚合
type ConcreteAggregate struct {
	elements []interface{}
}

func (a *ConcreteAggregate) CreateIterator() Iterator {
	return &ConcreteIterator{
		aggregate: a,
		index:     0, // 遍历是从0开始
		lastIndex: len(a.elements) - 1,
	}
}

// ConcreteIterator 具体迭代器
type ConcreteIterator struct {
	aggregate *ConcreteAggregate
	index     int
	lastIndex int
}

func (i *ConcreteIterator) HasNext() bool {
	return i.index < len(i.aggregate.elements)
}

func (i *ConcreteIterator) Next() interface{} {
	element := i.aggregate.elements[i.index]
	i.index++
	return element
}

func (i *ConcreteIterator) HasPrev() bool {
	return i.lastIndex >= 0

}
func (i *ConcreteIterator) Prev() interface{} {
	element := i.aggregate.elements[i.lastIndex]
	i.lastIndex--
	return element
}
