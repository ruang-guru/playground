package main

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestRotateCp(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "RotateCp Suite")
}
