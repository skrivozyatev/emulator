package emulator

import (
	"log"
	"math/rand"
	"os"
	"sync"
	"time"
)

type input struct {
	id     int
	sorter *sorter
	rnd    *rand.Rand
	info   *log.Logger
}

func CreateInput(id int, sorter *sorter) *input {
	inp := new(input)
	inp.id = id
	inp.rnd = rand.New(rand.NewSource(time.Now().UnixNano()))
	inp.info = log.New(os.Stdout, "[INFO] ", log.Ldate|log.Ltime)
	inp.sorter = sorter
	return inp
}

func (inp *input) Put(count int, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 0; i < count; i++ {
		wg.Add(1)
		parcel := inp.sorter.parcelFactory.GenParcel()
		inp.info.Printf("Put parcel %d into input %d", parcel.id, inp.id)
		go inp.sorter.carrier.Put(inp.id, parcel, wg)
		inp.sorter.carrier.sleep(2)
	}
}
