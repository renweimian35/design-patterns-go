// Package factory 工厂模式
package factory

type Car interface {
	message() string
}

type Benz struct {
}

func (b *Benz) message() string {
	return "梅赛德斯-奔驰（Mercedes-Benz）是世界闻名的德国豪华汽车品牌。"
}

type BMW struct {
}

func (b *BMW) message() string {
	return "宝马（BMW），中文全称为巴伐利亚发动机制造厂股份有限公司，德国汽车品牌。"
}

type Mazda struct {
}

func (m *Mazda) message() string {
	return "马自达(Mazda)，堵车最厉害的品牌"
}

type Nil struct {
}

func (m *Nil) message() string {
	return "找不到该品牌"
}

// NewCarFactory 具体工厂实现，用于创建具体汽车
func NewCarFactory(carName string) Car {
	switch carName {
	case "奔驰":
		return &Benz{}
	case "宝马":
		return &BMW{}
	case "马自达":
		return &Mazda{}
	default:
		return &Nil{}
	}
}
