package main

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Parentheses Matching", func() {
	Describe("Given empty parameter", func() {
		It("Should return false", func() {
			got := IsValidParentheses("")
			Expect(got).To(BeFalse())
		})
	})
	Describe("Given correct parentheses sequence", func() {
		It("Should return true", func() {
			got := IsValidParentheses("()")
			Expect(got).To(BeTrue())
		})
	})
	Describe("Given correct parentheses sequence", func() {
		It("Should return true", func() {
			got := IsValidParentheses("(){}[]")
			Expect(got).To(BeTrue())
		})
	})
	Describe("Given unbalanced parentheses sequence", func() {
		It("Should return false", func() {
			got := IsValidParentheses("{[()]")
			Expect(got).To(BeFalse())
		})
	})
	Describe("Given incorrect parentheses sequence", func() {
		It("Should return false", func() {
			got := IsValidParentheses("[{)]")
			Expect(got).To(BeFalse())
		})
	})
	Describe("Given only open parentheses sequence", func() {
		It("Should return false", func() {
			got := IsValidParentheses("({{")
			Expect(got).To(BeFalse())
		})
	})
	Describe("Given only close parentheses sequence", func() {
		It("Should return false", func() {
			got := IsValidParentheses("]])")
			Expect(got).To(BeFalse())
		})
	})
})
