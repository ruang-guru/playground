package main

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Main", func() {
	Describe("Delete Value", func() {
		When("there is node is not empty at linked list", func() {
			It("should return linked list after deleted value and nil error", func() {
				list := new(LinkedList)
				list.head = nil
				list.len = 0

				list.Insert(1)
				list.Insert(2)
				list.Insert(3)
				list.Insert(4)
				linkedlist, err := list.DeleteVal(3)
				Expect(err).To(BeNil())
				Expect(linkedlist.len).To(Equal(3))
				Expect(linkedlist.head.value).To(Equal(1))
				Expect(linkedlist.head.next.value).To(Equal(2))
				Expect(linkedlist.head.next.next.value).To(Equal(4))
			})
		})
	})
})
