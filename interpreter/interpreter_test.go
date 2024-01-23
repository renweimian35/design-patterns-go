package interpreter

import (
	"fmt"
	"testing"
)

func TestInterpreter(t *testing.T) {
	// 构建解释器
	parser := &Parser{}
	// 输入车辆配置描述
	description := "Red Bus"
	// 解析车辆配置
	parser.Parse(description)
	// 创建上下文
	context := &Context{car: &Car{}}
	// 解释并执行车辆配置
	parser.Interpret(context)

	// 打印结果
	fmt.Printf("Car Configuration: Model=%s,Color=%s\n",
		context.car.Model, context.car.Color)
}
