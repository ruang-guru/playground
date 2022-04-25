package main

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestServerCp(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "ServerCp Suite")
}
