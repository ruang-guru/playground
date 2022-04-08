package main

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Main", func() {
	Describe("MiddleNode", func() {
		When("there is one middle node", func() {
			It("should return middle node along with the next value", func() {
				node1 := new(ListNode)
				node1.Val = 1

				node2 := new(ListNode)
				node2.Val = 2
				node1.Next = node2

				node3 := new(ListNode)
				node3.Val = 3
				node2.Next = node3

				node4 := new(ListNode)
				node4.Val = 4
				node3.Next = node4

				node5 := new(ListNode)
				node5.Val = 5
				node4.Next = node5

				node := node1.MiddleNode(node1)
				Expect(node.Val).To(Equal(3))
				Expect(node.Next.Val).To(Equal(4))
				Expect(node.Next.Next.Val).To(Equal(5))
			})
		})
		When("there is two middle node", func() {
			It("should return the second middle node along with the next value", func() {
				node1 := new(ListNode)
				node1.Val = 1

				node2 := new(ListNode)
				node2.Val = 2
				node1.Next = node2

				node3 := new(ListNode)
				node3.Val = 3
				node2.Next = node3

				node4 := new(ListNode)
				node4.Val = 4
				node3.Next = node4

				node5 := new(ListNode)
				node5.Val = 5
				node4.Next = node5

				node6 := new(ListNode)
				node6.Val = 6
				node5.Next = node6

				node := node1.MiddleNode(node1)
				Expect(node.Val).To(Equal(4))
				Expect(node.Next.Val).To(Equal(5))
				Expect(node.Next.Next.Val).To(Equal(6))
			})
		})
	})
})
