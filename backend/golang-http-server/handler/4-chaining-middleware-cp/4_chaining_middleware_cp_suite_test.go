package main_test

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func Test4ChainingMiddlewareCp(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "4ChainingMiddlewareCp Suite")
}
