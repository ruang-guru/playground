package main

type Item struct {
	id    int
	name  string
	price float64
}

type Inventory struct {
	items []Item
}

func NewItem(id int, name string, price float64) Item {
	return Item{id, name, price}
}

func (i *Inventory) Add(item Item) {
	// TODO: answer here
}

func (i Inventory) Items() []Item {
	return i.items
}
