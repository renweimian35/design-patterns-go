package abstract_factory

import (
	"fmt"
	"testing"
)

func TestNewCar(t *testing.T) {
	factory := NewFactory()
	car := factory.CreateMazda()
	fmt.Println(car.message())
}
