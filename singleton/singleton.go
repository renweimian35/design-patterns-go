package singleton

import "sync"

// Singleton 单例类
type Singleton struct {
	car   string
	price int
	mux   sync.Mutex // 防止竞争，上锁
}

var instance *Singleton
var once sync.Once

// GetInstance 获取单例实例
func GetInstance() *Singleton {
	once.Do(func() {
		instance = &Singleton{}
	})
	return instance
}

func (s *Singleton) SetCar(name string) {
	s.mux.Lock()
	defer s.mux.Unlock()
	s.car = name
}

func (s *Singleton) SetPrice() {
	s.mux.Lock()
	defer s.mux.Unlock()
	s.price += 10000
}

func (s *Singleton) Message() (string, int) {
	s.mux.Lock()
	defer s.mux.Unlock()
	return s.car, s.price
}
