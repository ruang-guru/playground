package main

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func Test3AnonymousFunctionCp(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "3AnonymousFunctionCp Suite")
}
