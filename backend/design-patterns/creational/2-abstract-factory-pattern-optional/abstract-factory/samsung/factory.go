package samsung

import (
	abstract_factory "github.com/ruang-guru/playground/backend/design-patterns/creational/2-abstract-factory-pattern-optional/abstract-factory"
	"github.com/ruang-guru/playground/backend/design-patterns/creational/2-abstract-factory-pattern-optional/abstract-factory/samsung/product"
)

type Samsung struct {
}

func (s *Samsung) CreateSmartphone() abstract_factory.Smartphone {
	return &product.SamsungS22{}
}
func (s *Samsung) CreateTablet() abstract_factory.Tablet {
	return &product.SamsungTabS8{}
}
