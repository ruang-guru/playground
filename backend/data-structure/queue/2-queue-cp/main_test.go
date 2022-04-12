package main

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Main", func() {
	Describe("Insert Front", func() {
		When("first queue is not filled", func() {
			It("should insert front", func() {
				queues, err := q.InsertFront("Aditira")
				Expect(err).To(BeNil())
				Expect(queues).To(Equal(Queue{"Aditira", "nil", "nil", "nil", "nil"}))
			})
		})
		When("first queue is already filled", func() {
			It("should cannot insert front queue and return erron", func() {
				queues, err := q.InsertFront("Aditira")
				Expect(err.Error()).To(Equal("element cannot be inserted"))
				Expect(queues).To(Equal(Queue{"Aditira", "nil", "nil", "nil", "nil"}))
			})
		})
	})

	Describe("Insert Rear", func() {
		When("rear queue still has slots", func() {
			It("should insert rear", func() {
				q.InsertRear("Juno")
				q.InsertRear("Miya")
				q.InsertRear("Gina")
				queues, err := q.InsertRear("Pito")
				Expect(err).To(BeNil())
				Expect(queues).To(Equal(Queue{"Aditira", "Juno", "Miya", "Gina", "Pito"}))

			})
		})
		When("queue is full", func() {
			It("should return error", func() {
				_, err := q.InsertRear("Rian")
				Expect(err.Error()).To(Equal("Overflow"))
			})
		})
	})
})
