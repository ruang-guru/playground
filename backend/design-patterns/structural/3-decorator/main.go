package main

import (
	"fmt"

	"github.com/ruang-guru/playground/backend/design-patterns/structural/3-decorator/pintu"
)

func main() {
	door := &pintu.Pintu{}
	door.Open()

	fmt.Println("=========")

	smartDoor := pintu.NewSmartGagang(door)
	smartDoor.Open()
}
