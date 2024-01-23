package prototype

import (
	"fmt"
	"testing"
)

func TestPrototype(t *testing.T) {
	// 创建原型对象
	prototype := NewConcretePrototype("milo")

	// 克隆原型对象
	clone1 := prototype.Clone()

	// 源对象修改后不影响clone后的对象
	prototype.name = "miloyang"

	fmt.Println(prototype.name)
	fmt.Println(clone1.Name())

}
