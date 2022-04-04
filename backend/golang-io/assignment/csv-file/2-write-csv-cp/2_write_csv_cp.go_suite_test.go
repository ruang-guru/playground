package main_test

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func Test2WriteCsvCp(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "2WriteCsvCp.Go Suite")
}
