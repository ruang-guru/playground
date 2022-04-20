package main

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func Test5ChannelDirectionCp(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "5ChannelDirectionCp Suite")
}
