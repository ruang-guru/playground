package main

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestBinnarySearchCp(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "BinnarySearchCp Suite")
}
