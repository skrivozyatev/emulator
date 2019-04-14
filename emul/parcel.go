package emulator

import (
	"math/rand"
	"time"
)

type parcel struct {
	id       int64
	BarCodes string
	Length   int
	Width    int
	Height   int
	Weight   int
}

type parcelFactory struct {
	rnd *rand.Rand
}

func CreateParcelFactory() *parcelFactory {
	pf := new(parcelFactory)
	pf.rnd = rand.New(rand.NewSource(time.Now().UnixNano()))
	return pf
}

func DeclareParcelArray(n int) [](*parcel) {
	return make([](*parcel), n)
}

func (factory *parcelFactory) CreateParcel(barCodes string, length int, width int, height int, weight int) *parcel {
	p := new(parcel)
	p.id = factory.rnd.Int63()
	p.BarCodes = barCodes
	p.Length = length
	p.Width = width
	p.Height = height
	p.Weight = weight
	return p
}

func (factory *parcelFactory) GenParcel() *parcel {
	return factory.CreateParcel("12345;678910", factory.rnd.Intn(1000), factory.rnd.Intn(1000),
		factory.rnd.Intn(1000), factory.rnd.Intn(10000))
}
