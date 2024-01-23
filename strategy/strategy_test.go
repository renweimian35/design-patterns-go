package strategy

import "testing"

func TestStrategy(t *testing.T) {
	// 创建上下文对象
	context := &Context{}

	// 使用走国道策略
	nRoad := &NationalRoad{}
	context.SetStrategy(nRoad)
	context.ExecuteStrategy()

	// 使用走高速策略
	expy := &Expy{}
	context.SetStrategy(expy)
	context.ExecuteStrategy()
}
