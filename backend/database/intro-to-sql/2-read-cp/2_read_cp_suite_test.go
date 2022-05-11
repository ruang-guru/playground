package main

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func Test2ReadCp(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "2ReadCp Suite")
}
