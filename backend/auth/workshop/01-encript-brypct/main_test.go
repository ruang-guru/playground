package main

import (
	"fmt"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"golang.org/x/crypto/bcrypt"
)

var _ = Describe("Main", func() {
	Describe("encrypt to bcrypt", func() {
		It("should return hashed str based on request", func() {
			str := "hello this is key"

			hashedStr, err := encryptToBcrypt(str)
			if err != nil {
				return
			}
			Expect(err).To(BeNil())

			// Comparing the password with the hash
			err = bcrypt.CompareHashAndPassword([]byte(hashedStr), []byte(str))
			fmt.Println(err) // nil means it is a match

			Expect(err).To(BeNil())

		})
	})
})
