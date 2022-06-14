package api_test

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestAPICashierapp(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "APICashierapp Suite")
}
