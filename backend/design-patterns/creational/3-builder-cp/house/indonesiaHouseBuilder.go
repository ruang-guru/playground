package house

type indonesiaHouseBuilder struct {
	house House
}

// di Indonesia kita bebas tuh mau berapa window nya
func (i *indonesiaHouseBuilder) buildWindow(numOfWindow int) {
	i.house.NumOfWindows = numOfWindow
}

func (i *indonesiaHouseBuilder) buildDoor() {
	i.house.NumOfDoors++
}

func (i *indonesiaHouseBuilder) buildGarage() {
	i.house.HasGarage = true
}

func (i *indonesiaHouseBuilder) buildSwimmingPool() {
	i.house.HasSwimmingPool = true
}

func (i *indonesiaHouseBuilder) getHouse() House {
	return i.house
}
