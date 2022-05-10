package product

import abstract_factory "github.com/ruang-guru/playground/backend/design-patterns/creational/2-abstract-factory-pattern-optional/abstract-factory"

type SamsungTabS8 struct{}

func (l *SamsungTabS8) Price() float64 {
	return 17999999
}

func (l *SamsungTabS8) Is5G() bool {
	return true
}

func (l *SamsungTabS8) IsFoldable() bool {
	return false
}

func (l *SamsungTabS8) Size() abstract_factory.Dimension {
	return abstract_factory.Dimension{
		Length: 8,
		Width:  6,
		Height: 1,
	}
}
