package main

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func Test1ReadCsvCp(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "1ReadCsvCp Suite")
}
