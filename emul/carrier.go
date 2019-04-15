package emulator

import (
	"log"
	"math/rand"
	"sync"
	"time"
)

type carrier struct {
	rnd *rand.Rand
}

func CreateCarrier() *carrier {
	c := new(carrier)
	c.rnd = rand.New(rand.NewSource(time.Now().UnixNano()))
	return c
}

func (c *carrier) Put(parcel *parcel, wg *sync.WaitGroup) {
	go c.put(parcel, wg)
}

func (c *carrier) put(parcel *parcel, wg *sync.WaitGroup) {
	defer wg.Done()
	log.Printf("Parcel %d start\n", parcel.id)
	c.carry(1)
	log.Printf("Parcel %d request 1\n", parcel.id)
	c.carry(1)
	log.Printf("Parcel %d request 2\n", parcel.id)
	c.carry(1)
	log.Printf("Parcel %d request 3\n", parcel.id)
	c.carry(2)
	log.Printf("Parcel %d report\n", parcel.id)
}

func (c *carrier) carry(seconds int) {
	time.Sleep(time.Duration(seconds)*time.Second + time.Duration(c.rnd.Intn(1000))*time.Millisecond)
}
