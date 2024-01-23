package state

import "testing"

func TestState(t *testing.T) {
	car := &Car{}

	lock := &ChildLock{}
	car.SetState(lock) // 车辆上儿童锁
	car.Request()

	unlock := &ChildUnLock{}
	car.SetState(unlock) // 车辆解锁儿童锁
	car.Request()
}
