package main

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestCountDigitCp(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "CountDigitCp Suite")
}
