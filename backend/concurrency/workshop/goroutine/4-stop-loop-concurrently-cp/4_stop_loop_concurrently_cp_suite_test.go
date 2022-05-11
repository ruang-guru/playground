package main

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func Test4StopLoopConcurrentlyCp(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "4StopLoopConcurrentlyCp Suite")
}
