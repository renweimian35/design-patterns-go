package memento

import (
	"fmt"
	"testing"
)

func TestMemento(t *testing.T) {
	zhangsan := &Originator{State: "紧急", Item: "今天要去看马自达"}
	caretaker := &Caretaker{}

	caretaker.Memento = zhangsan.CreateMemento() // 把zhangsan的备忘录加到管理者中去

	fmt.Println(zhangsan)

	zhangsan.State = "不太紧急"
	zhangsan.Item = "今天去看宝马"

	// 不行，我还是恢复到原来的备忘吧
	zhangsan.SetMemento(caretaker.Memento)

	fmt.Println(zhangsan)

}
