// Package observer 观察者模式
package observer

import (
	"fmt"
	"sync"
)

// CarObserver 观察者接口
type CarObserver interface {
	breakdown(level int, info string)
}

// CarInfoCenter 主题接口
type CarInfoCenter interface {
	RegisterObserver(observer CarObserver)
	RemoveObserver(observer CarObserver)
	NotifyObservers()
	SetCarStatus(level int, info string)
}

// CarInfoCenterImpl 具体主题
type CarInfoCenterImpl struct {
	observers []CarObserver
	mutex     sync.Mutex
	info      string
	level     int
}

func (cic *CarInfoCenterImpl) RegisterObserver(observer CarObserver) {
	cic.mutex.Lock()
	defer cic.mutex.Unlock()
	cic.observers = append(cic.observers, observer)
}
func (cic *CarInfoCenterImpl) RemoveObserver(observer CarObserver) {
	cic.mutex.Lock()
	defer cic.mutex.Unlock()
	for i, carObserver := range cic.observers {
		if carObserver == observer {
			cic.observers = append(cic.observers[:i], cic.observers[i+1:]...)
			break
		}
	}
}

func (cic *CarInfoCenterImpl) NotifyObservers() {
	cic.mutex.Lock()
	defer cic.mutex.Unlock()
	for _, observer := range cic.observers {
		observer.breakdown(cic.level, cic.info)
	}
}

func (cic *CarInfoCenterImpl) SetCarStatus(level int, info string) {
	cic.info = info
	cic.level = level
	cic.NotifyObservers()
}

type CarDisplay struct {
}

func (c *CarDisplay) breakdown(level int, info string) {
	fmt.Println("car display receive breakdown level:", level, " , info:", info)
}

type AppDisplay struct {
}

func (a *AppDisplay) breakdown(level int, info string) {
	fmt.Println("app display receive breakdown level:", level, " , info:", info)
}
