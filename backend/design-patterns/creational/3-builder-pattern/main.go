package main

import (
	"fmt"

	"github.com/ruang-guru/playground/backend/design-patterns/creational/3-builder-pattern/house"
)

func main() {
	indonesianHouseBuilder := house.NewBuilder("indonesia")
	zimbabweHouseBuilder := house.NewBuilder("zimbabwe")

	kontraktor := house.NewKontraktor(indonesianHouseBuilder)
	indonesianHouse := kontraktor.BuildHouse()

	fmt.Printf("Rumah di Indonesia mempunyai jumlah jendela sebanyak: %d\n", indonesianHouse.NumOfWindows)
	fmt.Printf("Rumah di Indonesia mempunyai jumlah pintu sebanyak: %d\n", indonesianHouse.NumOfDoors)
	fmt.Printf("Rumah di Indonesia mempunyai garasi: %t\n", indonesianHouse.HasGarage)
	fmt.Printf("Rumah di Indonesia mempunyai kolam berenang: %t\n", indonesianHouse.HasSwimmingPool)

	kontraktor = house.NewKontraktor(zimbabweHouseBuilder)
	zimbabweHouse := kontraktor.BuildHouse()

	fmt.Printf("Rumah di Zimbabwe mempunyai jumlah jendela sebanyak: %d\n", zimbabweHouse.NumOfWindows)
	fmt.Printf("Rumah di Zimbabwe mempunyai jumlah pintu sebanyak: %d\n", zimbabweHouse.NumOfDoors)
	fmt.Printf("Rumah di Zimbabwe mempunyai garasi: %t\n", zimbabweHouse.HasGarage)
	fmt.Printf("Rumah di Zimbabwe mempunyai kolam berenang: %t\n", zimbabweHouse.HasSwimmingPool)

	zimbabweWithoutSimmingPool := house.NewKontraktor(zimbabweHouseBuilder).BuildHouseWithoutSwimmingPool()

	fmt.Printf("Rumah di Zimbabwe mempunyai jumlah jendela sebanyak: %d\n", zimbabweWithoutSimmingPool.NumOfWindows)
	fmt.Printf("Rumah di Zimbabwe mempunyai jumlah pintu sebanyak: %d\n", zimbabweWithoutSimmingPool.NumOfDoors)
	fmt.Printf("Rumah di Zimbabwe mempunyai garasi: %t\n", zimbabweWithoutSimmingPool.HasGarage)
	fmt.Printf("Rumah di Zimbabwe mempunyai kolam berenang: %t\n", zimbabweWithoutSimmingPool.HasSwimmingPool)
}
