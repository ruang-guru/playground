package main

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func Test5Close2Cp(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "5Close2Cp Suite")
}
