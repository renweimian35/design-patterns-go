// Package facade 外观模式
package facade

type operate interface {
	open() string
	close() string
}

type engine struct {
}

func (e *engine) open() string {
	return "open engine"
}

func (e *engine) close() string {
	return "close engine"
}

type music struct {
}

func (e *music) open() string {
	return "open music"
}

func (e *music) close() string {
	return "close music"
}

type Facade struct {
	eng engine
	mus music
}

func NewFacade() *Facade {
	return &Facade{
		eng: engine{},
		mus: music{},
	}
}

// 启动车子
func (f *Facade) startCar() string {
	result := "车子启动中...\n"
	result += f.eng.open() + "\n"
	result += f.mus.open()
	return result
}

// 车子熄火
func (f *Facade) stopCar() string {
	result := "车子熄火中...\n"
	result += f.eng.close() + "\n"
	result += f.mus.close()
	return result
}
