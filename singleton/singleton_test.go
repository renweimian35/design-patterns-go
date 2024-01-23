package singleton

import (
	"fmt"
	"sync"
	"testing"
)

func TestSingleton(t *testing.T) {
	var wg sync.WaitGroup
	wg.Add(200)

	for i := 0; i < 200; i++ {
		go func() {
			instance := GetInstance()
			instance.SetCar("BWM")
			instance.SetPrice()
			wg.Done()
		}()
	}

	wg.Wait()
	instance1 := GetInstance()
	fmt.Println(instance1.Message()) // BWM 200
}
