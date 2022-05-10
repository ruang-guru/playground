package house

type zimbabweHouseBuilder struct {
	house House
}

// di Zimbabwe sendiri kita hanya bisa membuat jendela dengan maksimal 2
func (i *zimbabweHouseBuilder) buildWindow(numOfWindow int) {
	if numOfWindow > 2 {
		i.house.NumOfWindows = 2
	} else {
		i.house.NumOfWindows = numOfWindow
	}
}

func (i *zimbabweHouseBuilder) buildDoor() {
	i.house.NumOfDoors++
}

func (i *zimbabweHouseBuilder) buildGarage() {
	i.house.HasGarage = true
}

func (i *zimbabweHouseBuilder) buildSwimmingPool() {
	if i.house.HasSwimmingPool {
		i.house.HasSwimmingPool = false
	} else {
		i.house.HasSwimmingPool = true
	}
}

func (i *zimbabweHouseBuilder) getHouse() House {
	return i.house
}
