package main

import "fmt"

type PrimaryKey int

type WarehouseStockRow struct {
	ID       PrimaryKey
	Code     string
	ItemName string
	Unit     string
	Quantity int
	Price    int
}

type WarehouseDB struct {
	m map[PrimaryKey]WarehouseStockRow
}

func main() {
	db := NewWarehouse()
	db.Insert("B001", "Aqua", "Btl", 22, 2000)
	db.Insert("B002", "Coca Cola", "Btl", 2, 1000)
	db.Insert("B003", "Fanta", "Btl", 5, 5000)
	db.Insert("B004", "Teh Botol", "Kotak", 35, 2500)

	db.Delete(3)

	fmt.Println(db.m)
}

func NewWarehouse() *WarehouseDB {
	return &WarehouseDB{
		m: make(map[PrimaryKey]WarehouseStockRow),
	}
}

func (db *WarehouseDB) Insert(code string, itemName string, unit string, quantity int, price int) {
	db.m[PrimaryKey(len(db.m)+1)] = WarehouseStockRow{
		ID:       PrimaryKey(len(db.m)) + 1,
		Code:     code,
		ItemName: itemName,
		Unit:     unit,
		Quantity: quantity,
		Price:    price,
	}
}

func (db *WarehouseDB) Delete(id PrimaryKey) {
	delete(db.m, id)
}
