// Package template 模板方法
package template

import "fmt"

// AbstractDrive 模板类
type AbstractDrive interface {
	OpenDoor()
	Fire()
	Going()
	Hook()
}

type Driver struct {
	AbstractDrive
}

func NewDriver(abstract AbstractDrive) *Driver {
	return &Driver{
		abstract,
	}
}

func (d *Driver) Template() {
	d.OpenDoor()
	d.Fire()
	d.Going()
	d.Hook()
}

// CommonTemplate 具体模板类
type CommonTemplate struct {
}

func (d *CommonTemplate) OpenDoor() {
	fmt.Println("开门")
}

func (d *CommonTemplate) Fire() {
	fmt.Println("点火")
}

func (d *CommonTemplate) Going() {
	fmt.Println("起步")
}

// Hook 钩子方法，子类可以选择性地覆盖
func (d *CommonTemplate) Hook() {
	fmt.Println("开启音乐")
}
