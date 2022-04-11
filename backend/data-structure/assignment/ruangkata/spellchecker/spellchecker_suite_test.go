package spellchecker_test

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestSpellchecker(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Spellchecker Suite")
}
