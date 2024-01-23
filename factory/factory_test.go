package factory

import (
	"fmt"
	"testing"
)

func TestNewCar(t *testing.T) {
	message := NewCarFactory("马自达").message()
	fmt.Println(message)
}
