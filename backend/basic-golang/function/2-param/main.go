package main

import "fmt"

func main() {
	hello("teman")
	hello("teman 2")

}

func hello(name string) {
	fmt.Println("Halo", name)
}
