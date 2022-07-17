package apple

import (
	abstract_factory "github.com/ruang-guru/playground/backend/design-patterns/creational/2-abstract-factory-pattern-optional/abstract-factory"
	"github.com/ruang-guru/playground/backend/design-patterns/creational/2-abstract-factory-pattern-optional/abstract-factory/apple/product"
)

type Apple struct {
}

func (s *Apple) CreateSmartphone() abstract_factory.Smartphone {
	return &product.IPhone13{}
}
func (s *Apple) CreateTablet() abstract_factory.Tablet {
	return &product.IpadAir4{}
}
