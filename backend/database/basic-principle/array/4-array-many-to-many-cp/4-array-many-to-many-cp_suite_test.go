package main_test

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func Test4ArrayManyToManyCp(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "4ArrayManyToManyCp Suite")
}
