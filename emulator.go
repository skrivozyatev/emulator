package main

import (
	"./config"
	"./emul"
	"sync"
)

func main() {
	wg := new(sync.WaitGroup)
	sorter := emulator.CreateSorter()
	for i := 0; i < config.GetInputCount(); i++ {
		inp := emulator.CreateInput(i+1, sorter)
		wg.Add(1)
		go inp.Put(config.GetParcelsPerInput(), wg)
	}
	wg.Wait()
}
