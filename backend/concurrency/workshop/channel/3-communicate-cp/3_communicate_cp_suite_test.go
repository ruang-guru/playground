package main

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func Test3CommunicateCp(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "3CommunicateCp Suite")
}
