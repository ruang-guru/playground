package main_test

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestUnittest(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Unittest Suite")
}
