package main

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestFind2NumberAddsUptoTarget(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Find2NumberAddsUptoTarget")
}
