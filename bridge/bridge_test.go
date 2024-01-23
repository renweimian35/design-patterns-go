package bridge

import (
	"fmt"
	"testing"
)

func TestBridge(t *testing.T) {
	red := ColorRed{}
	large := Large{}
	large.SetColorImplementor(red)
	fmt.Println(large.Hub()) // hub size large:8.0J,color:Red

	blue := ColorBlue{}
	medium := Medium{}
	medium.SetColorImplementor(blue)
	fmt.Println(medium.Hub()) // hub size medium:7.0J,color:Blue
}
