package main

import "fmt"

type DriverRow struct {
	ID     int // primary key
	Name   string
	Age    int
	Region string
	Car    CarRow
}

type CarRow struct {
	ID    int // primary key
	Name  string
	Model string
	Maker string
	Price int
}

type DriverTable []DriverRow

func main() {
	db := make(DriverTable, 0)

	db.InsertDrivers("John", 21, "Jakarta", CarRow{ID: 1, Name: "Honda City", Model: "City", Maker: "Honda", Price: 20000})
	db.InsertDrivers("Jane", 22, "Bandung", CarRow{ID: 1, Name: "Honda City", Model: "City", Maker: "Honda", Price: 20000})
	db.InsertDrivers("Gina", 24, "Jakarta", CarRow{ID: 1, Name: "Honda City", Model: "City", Maker: "Honda", Price: 20000})
	db.InsertDrivers("Vina", 25, "Cianjur", CarRow{ID: 2, Name: "Bugatti Chiron", Model: "Chiron", Maker: "Bugatti", Price: 3000000})

	for _, row := range db {
		db.String(row)
	}
}

func (db *DriverTable) String(row DriverRow) {
	fmt.Printf("| ID: %d | Driver: %v | Age: %d | Region: %v | Car: %v |\n", row.ID, row.Name, row.Age, row.Region, row.Car)
}

func (db *DriverTable) InsertDrivers(name string, age int, region string, car CarRow) {
	(*db) = append(*db, DriverRow{
		ID:     len(*db) + 1,
		Name:   name,
		Age:    age,
		Region: region,
		Car:    car,
	})
}
