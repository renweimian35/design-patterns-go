package iterator

import (
	"fmt"
	"testing"
)

func TestIterator(t *testing.T) {
	aggregate := &ConcreteAggregate{
		elements: []interface{}{"BMW", "Benz", "Mazda", "BYD", 1, 2, 3, 4},
	}

	iterator := aggregate.CreateIterator()

	for iterator.HasNext() {
		element := iterator.Next()
		fmt.Println(element)
	}

	fmt.Println("=================")

	for iterator.HasPrev() {
		element := iterator.Prev()
		fmt.Println(element)

	}
}
