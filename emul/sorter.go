package emulator

type sorter struct {
	carrier       *carrier
	parcelFactory *parcelFactory
}

func CreateSorter() *sorter {
	sorter := new(sorter)
	sorter.carrier = CreateCarrier()
	sorter.parcelFactory = CreateParcelFactory()
	return sorter
}
