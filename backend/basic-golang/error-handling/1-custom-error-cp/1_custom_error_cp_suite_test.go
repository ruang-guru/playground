package main_test

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func Test1CustomErrorCp(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "1CustomErrorCp Suite")
}
