package proxy

import "testing"

func TestProxy(t *testing.T) {
	agent := Agent{
		boss: Boss{},
	}
	agent.buyCar("BWM", 100000) //委托中介买车，车品牌是BWM,价格是100000,中介费是50000--- PASS: TestPro
}
