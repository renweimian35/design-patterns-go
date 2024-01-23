package adapter

import "fmt"

// USB 目标接口
type USB interface {
	Charge()
}

// CigaretteLighter 被适配者，点烟孔
type CigaretteLighter struct {
	ConnectedDevice string
}

func (c *CigaretteLighter) PowerSupply() string {
	return fmt.Sprintf("Power supply from cigarette lighter to %s", c.ConnectedDevice)
}

// Adapter 适配器 那一坨东西
type Adapter struct {
	CigaretteLighter *CigaretteLighter // 它需要插在点烟孔上面
}

// Charge 转成USB接口
func (adapter *Adapter) Charge() {
	fmt.Println(adapter.CigaretteLighter.PowerSupply())
}
