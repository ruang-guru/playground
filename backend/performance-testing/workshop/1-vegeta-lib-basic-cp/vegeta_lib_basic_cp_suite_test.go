package main

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestVegetaLibBasicCp(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "VegetaLibBasicCp Suite")
}
