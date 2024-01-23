package responsibility_chain

import (
	"fmt"
	"testing"
)

func TestResponsibilityChain(t *testing.T) {
	// 客户端使用责任链模式
	bmwHandler := &CarHandler{Price: 200000, Car: "BMW"}
	benzHandler := &CarHandler{Price: 150000, Car: "Benz"}
	mazdaHandler := &CarHandler{Price: 100000, Car: "Mazda"}
	noCarHandler := &CarHandler{Price: 0, Car: "no car"}

	// 设置处理者的后继关系
	bmwHandler.SetSuccessor(benzHandler)
	benzHandler.SetSuccessor(mazdaHandler)
	mazdaHandler.SetSuccessor(noCarHandler)

	// 模拟用户输入不同的金额
	amounts := []int{180000, 220000, 120000, 80000}

	for _, amount := range amounts {
		// 通过责任链传递请求
		bmwHandler.HandleRequest(amount)
		fmt.Println("---------------------")
	}
}
