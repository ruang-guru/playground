package product

import abstract_factory "github.com/ruang-guru/playground/backend/design-patterns/creational/2-abstract-factory-pattern-optional/abstract-factory"

type IpadAir4 struct{}

func (l *IpadAir4) Price() float64 {
	return 17999999
}

func (l *IpadAir4) Is5G() bool {
	return false
}

func (l *IpadAir4) IsFoldable() bool {
	return false
}

func (l *IpadAir4) Size() abstract_factory.Dimension {
	return abstract_factory.Dimension{
		Length: 10,
		Width:  10,
		Height: 2,
	}
}
