package main

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func Test1QueueCp(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "1QueueCp Suite")
}
