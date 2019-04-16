package main

import (
	"./emul"
	"sync"
)

func main() {
	wg := new(sync.WaitGroup)
	sorter := emulator.CreateSorter()
	for i := 0; i < 6; i++ {
		inp := emulator.CreateInput(i+1, sorter)
		wg.Add(1)
		go inp.Put(20, wg)
	}
	wg.Wait()
}
