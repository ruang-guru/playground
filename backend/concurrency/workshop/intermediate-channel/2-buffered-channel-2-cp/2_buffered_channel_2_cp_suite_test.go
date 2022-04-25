package main

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func Test2BufferedChannel2Cp(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "2BufferedChannel2Cp Suite")
}
