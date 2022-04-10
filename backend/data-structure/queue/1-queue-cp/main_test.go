package main

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Main", func() {
	Describe("CountStudents", func() {
		When("the student queue and sandwich scores are the same", func() {
			It("students will take it and so that all students can eat, should return number of queues to 0", func() {
				var students = []int{1, 1, 1, 0, 0, 0}
				var sandwiches = []int{1, 1, 1, 0, 0, 0}
				Expect(CountStudents(students, sandwiches)).To(Equal(0))
			})
		})
		When("value of the student queue and and sanwitch is different, the student will return to the back queue", func() {
			It("continues until it reaches a point where all students want a sanwitch with a value of 1 but the topmost sanwitch stack is 0, should return remaining student queues to 3", func() {
				var students = []int{1, 1, 1, 0, 0, 1}
				var sandwiches = []int{1, 0, 0, 0, 1, 1}
				Expect(CountStudents(students, sandwiches)).To(Equal(3))
			})
		})
	})
})
