package house_test

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestHouseTest(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "HouseTest Suite")
}
