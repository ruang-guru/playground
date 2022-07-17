package main

import "time"

type Item struct {
	Title    string
	Deadline time.Time
}

type Todos struct {
	items []Item
}

func (todos *Todos) Add(item Item) {
	// TODO: answer here
}

func (todos *Todos) GetAll() []Item {
	return []Item{} // TODO: replace this
}

func (todos *Todos) GetUpcoming() []Item {
	return []Item{} // TODO: replace this
}

func NewItem(title string, deadline time.Time) Item {
	return Item{title, deadline}
}

func NewTodos() Todos {
	return Todos{}
}
