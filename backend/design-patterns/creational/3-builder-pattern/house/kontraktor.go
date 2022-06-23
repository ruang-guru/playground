package house

// nah ketika semua sudah kita implementasikan kita bisa membuat Kontraktor dimana kontraktor tersebut yang akan menentukan builder nya
type kontraktor struct {
	builder HouseBuilder
}

func NewKontraktor(builder HouseBuilder) *kontraktor {
	return &kontraktor{
		builder: builder,
	}
}

func (d *kontraktor) BuildHouse() House {
	d.builder.buildWindow(5)
	d.builder.buildDoor()
	d.builder.buildGarage()
	d.builder.buildSwimmingPool()
	return d.builder.getHouse()
}

func (d *kontraktor) BuildHouseWithoutSwimmingPool() House {
	d.builder.buildWindow(1)
	d.builder.buildDoor()
	d.builder.buildGarage()
	d.builder.buildSwimmingPool() // We need this. Because if we don't call this method, the house will have a swimming pool
	return d.builder.getHouse()
}
