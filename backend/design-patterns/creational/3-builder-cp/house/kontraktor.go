package house

// nah ketika semua sudah kita implementasikan kita bisa membuat Kontraktor dimana kontraktor tersebut yang akan menentukan builder nya
type kontraktor struct {
	builder HouseBuilder
}

func NewKontraktor(builder HouseBuilder) *kontraktor {
	return &kontraktor{
		// TODO: answer here
	}
}

func (d *kontraktor) BuildHouse() House {
	// TODO: answer here
	return d.builder.getHouse()
}

// Rumah yang tidak memiliki SwimmingPool hanya boleh mimiliki satu jendela
func (d *kontraktor) BuildHouseWithoutSwimmingPool() House {
	// TODO: answer here
	return d.builder.getHouse()
}
