package main

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func Test6LeftJoinCp(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "6LeftJoinCp Suite")
}
