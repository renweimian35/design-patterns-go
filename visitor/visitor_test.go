package visitor

import "testing"

func TestVisitor(t *testing.T) {
	e := new(Element)
	e.Accept(new(BMWVisitor))   // 我是宝马销售,我们的车是豪华品牌
	e.Accept(new(MazdaVisitor)) // 我是马自达销售,我们的车，现在不堵了
}
