package command

import "testing"

func TestCommand(t *testing.T) {
	// 客户端使用命令模式
	car := &CarSystem{}

	// 创建具体命令对象
	lockCommand := &LockCommand{Car: car}
	unlockCommand := &UnLockCommand{Car: car}

	// 创建调用者 车主
	carOwner := &CarOwner{}

	// 拿出钥匙，找到关锁按键
	carOwner.Command = lockCommand
	carOwner.PressButton() // 按下关锁按键

	carOwner.Command = unlockCommand
	carOwner.PressButton() // 按下开锁键
}
