package facade

import (
	"fmt"
	"testing"
)

func TestOperateCar(t *testing.T) {
	fmt.Println(NewFacade().startCar())
	fmt.Println(NewFacade().stopCar())
}
