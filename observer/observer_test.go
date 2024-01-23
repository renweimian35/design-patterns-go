package observer

import "testing"

func TestObserver(t *testing.T) {

	// 创建具体主题
	carInfoCenter := &CarInfoCenterImpl{}

	// 创建具体观察者
	carDisplay := &CarDisplay{}
	appDisplay := &AppDisplay{}

	// 注册观察者
	carInfoCenter.RegisterObserver(carDisplay)
	carInfoCenter.RegisterObserver(appDisplay)

	carInfoCenter.SetCarStatus(1, "胎压低，报警") // carDisplay和appDisplay都会收到通知

	carInfoCenter.RemoveObserver(appDisplay) // 移除app监听

	carInfoCenter.SetCarStatus(2, "燃油不足") // 只有carDisplay收到通知
}
