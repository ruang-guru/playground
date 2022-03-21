package main

import "fmt"

func main() {
	hello("adi", "jakarta")
	hello("ado", "surabaya")

}

func hello(name, home_town string) {
	fmt.Printf("Halo %s dari %s\n", name, home_town)
}
