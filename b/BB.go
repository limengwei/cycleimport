package b

import (
	"cycleimport/er"
	"fmt"
)

type B struct {
	a er.A
}

func (b *B) Hello() {
	fmt.Println(b.a.GetMsg())
}

func (b *B) GetMsg() string {
	return "hello AA"
}

func (b *B) SetA(a er.A) {
	b.a = a
}
