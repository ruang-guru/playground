package main

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func Test1ReadTextFileCp(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "1ReadTextFileCp Suite")
}
