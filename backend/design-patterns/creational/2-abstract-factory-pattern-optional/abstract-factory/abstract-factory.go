package abstract_factory

// nah disini kita membuatkan dulu interface interface yang akan kita gunakan

type Pricey interface {
	Price() float64
}

type Smartphone interface {
	Pricey
	Is5G() bool
	IsFoldable() bool
}

type Dimension struct {
	Length, Width, Height int
}

type Tablet interface {
	Pricey
	Size() Dimension
	IsFoldable() bool
}

type DeviceFactory interface {
	CreateSmartphone() Smartphone
	CreateTablet() Tablet
}
