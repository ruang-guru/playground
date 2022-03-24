package main_test

import (
	"time"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	todoapp "github.com/ruang-guru/playground/backend/basic-golang/todoapp"
)

var _ = Describe("Todo", func() {
	Describe("Add", func() {
		It("adds item to the todo list", func() {
			todos := todoapp.NewTodos()
			todos.Add(todoapp.NewItem(1, "Buy milk", time.Now()))
			Expect(todos.GetItems()).To(HaveLen(1))
		})
	})
})
