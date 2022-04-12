package main

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func Test1SetCp(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "1SetCp Suite")
}
