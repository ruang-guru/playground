package main_test

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func Test2CustomHandlerCp(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "2CustomHandlerCp Suite")
}
