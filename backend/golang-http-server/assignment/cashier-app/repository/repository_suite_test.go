package repository_test

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestCashierapp(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Cashierapp Suite")
}
