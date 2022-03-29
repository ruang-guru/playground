package main_test

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func Test1EncodeJsonCp(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "1EncodeJsonCp Suite")
}
