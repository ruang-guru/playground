package main

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Account", func() {
	Describe("ExecuteToByteBuffer", func() {
		It("returns slice of bytes", func() {
			account := Account{
				Name:    "Tony",
				Number:  "1002321",
				Balance: 1000,
			}
			b, err := ExecuteToByteBuffer(account)
			Expect(err).ShouldNot(HaveOccurred())
			Expect(string(b)).To(Equal("Akun Tony dengan Nomor 1002321 memiliki saldo sebesar $1000."))
		})
	})
})
