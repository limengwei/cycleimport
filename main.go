package main

import (
	"cycleimport/a"
	"cycleimport/b"
)

func main() {
	aa := new(a.A)
	bb := new(b.B)

	aa.SetB(bb)
	bb.SetA(aa)

	aa.Hello()

	bb.Hello()
}
