package main_test

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func Test1ServerApiPostCp(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Test1ServerApiPostCp Suite")
}
