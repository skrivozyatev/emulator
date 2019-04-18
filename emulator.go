package main

import (
	"./config"
	"./emul"
	"./logger"
	"sync"
	"time"
)

func main() {
	now := time.Now()
	wg := new(sync.WaitGroup)
	sorter := emulator.CreateSorter()
	emulator.StartReceiver(wg)
	for i := 0; i < config.InputCount; i++ {
		inp := emulator.CreateInput(i+1, sorter)
		wg.Add(1)
		go inp.Put(config.ParcelsPerInput, wg)
	}
	wg.Wait()
	logger.Info("Waiting for remaining replies...")
	time.Sleep(2 * time.Second)
	logger.Info("Done in", time.Since(now))
}
