package main_test

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func Test4ErrorIsErrorAsCp(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "4ErrorIsErrorAsCp Suite")
}
