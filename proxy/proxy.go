package proxy

import "fmt"

// Subject 主题接口，要买车
type Subject interface {
	buyCar(name string, price int)
}

// Boss 真实主题，老板要买车
type Boss struct {
}

func (b Boss) buyCar(name string, price int) {
	fmt.Printf("委托中介买车，车品牌是%s,价格是%d,中介费是%d", name, price, price/2)
}

// Agent 代理商
type Agent struct {
	boss Boss // 他代表老板
}

// 代理去买车
func (a *Agent) buyCar(name string, price int) {
	a.boss.buyCar(name, price) // 买完车通知老板
}
