package main

import (
	"./emul"
	"fmt"
	"sync"
)

func main() {
	pf := emulator.CreateParcelFactory()
	parcelCount := 6
	p := emulator.DeclareParcelArray(parcelCount)
	carrier := emulator.CreateCarrier()
	var wg sync.WaitGroup
	for i := 0; i < parcelCount; i++ {
		p[i] = pf.GenParcel()
		wg.Add(1)
		carrier.Put(p[i], &wg)
	}
	for i := 0; i < parcelCount; i++ {
		fmt.Println(p[i])
	}
	wg.Wait()
}
