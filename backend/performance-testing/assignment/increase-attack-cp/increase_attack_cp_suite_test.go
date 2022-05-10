package main

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestIncreaseAttackCp(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "IncreaseAttackCp Suite")
}
