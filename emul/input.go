package emulator

import (
	"../config"
	"../logger"
	"github.com/spf13/cast"
	"math/rand"
	"sync"
	"time"
)

type input struct {
	id     int
	sorter *sorter
	rnd    *rand.Rand
}

func CreateInput(id int, sorter *sorter) *input {
	inp := new(input)
	inp.id = id
	inp.rnd = rand.New(rand.NewSource(time.Now().UnixNano()))
	inp.sorter = sorter
	return inp
}

func (inp *input) Put(count int, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 0; i < count; i++ {
		wg.Add(1)
		parcel := inp.sorter.parcelFactory.GenParcel()
		logger.Infof("Put parcel %d into input %d", parcel.id, inp.id)
		go inp.sorter.carrier.Put(inp.id, parcel, wg)
		inp.sorter.carrier.sleep(cast.ToDuration(config.ParcelInputInterval))
	}
}
