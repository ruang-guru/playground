package main

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/ruang-guru/playground/backend/basic-golang/package/1-create-package-cp/account"
)

var _ = Describe("Account", func() {
	Describe("Get Balance", func() {
		It("returns 100 when account balance is 100", func() {
			acc := account.Account{
				Name:    "MyAccount",
				Balance: 100,
			}
			Expect(acc.GetBalance()).To(Equal(100))
		})
	})
	Describe("Deposit", func() {
		It("add account balance 100 if deposit 100", func() {
			acc := account.Account{
				Name:    "MyAccount",
				Balance: 100,
			}
			acc.Deposit(100)
			Expect(acc.GetBalance()).To(Equal(200))
		})
	})
})
