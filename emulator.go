package main

import (
	"./emul"
	"fmt"
)

func main() {
	pf := emulator.CreateParcelFactory()
	p := emulator.DeclareParcelArray(10)
	for i := 0; i < 10; i++ {
		p[i] = pf.GenParcel()
	}
	for i := 0; i < 10; i++ {
		fmt.Println(p[i])
	}
}
