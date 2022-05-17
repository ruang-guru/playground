package main_test

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func Test3ArrayOneToManyCp(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "3ArrayOneToManyCp Suite")
}
