// Package abstract_factory 普通工厂模式
package abstract_factory

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

// CarFactory 汽车工厂类，把所有的汽车全部在这里实现
type CarFactory interface {
	CreateBenz() Car
	CreateBMW() Car
	CreateMazda() Car
}

// simpleCarFactory 具体工厂的实现，用于创建产品族的产品
type simpleCarFactory struct {
}

func NewFactory() simpleCarFactory {
	return simpleCarFactory{}
}

func (s *simpleCarFactory) CreateBenz() Car {
	return &Benz{}
}

func (s *simpleCarFactory) CreateBMW() Car {
	return &BMW{}
}

func (s *simpleCarFactory) CreateMazda() Car {
	return &Mazda{}
}
