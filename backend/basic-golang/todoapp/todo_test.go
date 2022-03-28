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
			todos.Add(todoapp.NewItem("Buy milk", time.Now()))
			Expect(todos.GetAll()).To(HaveLen(1))
		})
	})

	Describe("Get Upcoming", func() {
		It("only returns upcoming todo items", func() {
			todos := todoapp.NewTodos()
			todos.Add(todoapp.NewItem("Buy milk", time.Now().Add(-time.Hour)))
			todos.Add(todoapp.NewItem("Buy eggs", time.Now().Add(time.Hour)))
			Expect(todos.GetUpcoming()).To(HaveLen(1))
			Expect(todos.GetUpcoming()[0].Title).To(Equal("Buy eggs"))
		})
	})

	Describe("Get All", func() {
		It("returns all todo items", func() {
			todos := todoapp.NewTodos()
			todos.Add(todoapp.NewItem("Buy milk", time.Now().Add(-time.Hour)))
			todos.Add(todoapp.NewItem("Buy eggs", time.Now().Add(time.Hour)))
			Expect(todos.GetAll()).To(HaveLen(2))
			Expect(todos.GetAll()[0].Title).To(Equal("Buy milk"))
			Expect(todos.GetAll()[1].Title).To(Equal("Buy eggs"))
		})
	})
})
