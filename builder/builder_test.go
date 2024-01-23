package builder

import (
	"fmt"
	"testing"
)

func TestBuilderCar(t *testing.T) {
	builder := &ConcreteCar{}
	director := Director{
		builder: builder,
	}
	car := director.Construct() // 车子组装了

	//打印车子的各种信息
	fmt.Println(car.Wheel)
	fmt.Println(car.Seat)
	fmt.Println(car.Steering)

}
