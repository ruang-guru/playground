package main

import "time"

type Item struct {
	id       int
	action   string
	deadline time.Time
}

type Todos struct {
	items []Item
}

func (todos *Todos) Add(item Item) {
	// TODO: answer here
}

func (todos *Todos) GetItems() []Item {
	// TODO: answer here
}

func NewItem(id int, action string, deadline time.Time) Item {
	return Item{id, action, deadline}
}

func NewTodos() Todos {
	// TODO: answer here
}
