package a

import (
	"cycleimport/er"
	"fmt"
)

type A struct {
	b er.B
}

func (a *A) Hello() {
	fmt.Println(a.b.GetMsg())
}

func (a *A) Hellooo() {
	fmt.Println("Hellooo")
}

func (a *A) GetMsg() string {
	return "hello BB"
}

func (a *A) SetB(b er.B) {
	a.b = b
}
