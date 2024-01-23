package decorator

import (
	"fmt"
	"testing"
)

func TestDecorator(t *testing.T) {
	basicCar := &BasicCar{}
	navigation := &NavigationDecorator{car: basicCar} // 加装导航系统
	fmt.Println(navigation.Drive())                   // NavigationDecorator: BasicCar: Driving with Navigation System
}
