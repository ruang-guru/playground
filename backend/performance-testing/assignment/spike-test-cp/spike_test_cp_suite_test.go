package main

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestSpikeTestCp(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "SpikeTestCp Suite")
}
