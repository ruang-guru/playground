package main_test

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func Test5RouteGroupingCp(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "5RouteGroupingCp Suite")
}
