package main

import (
	"encoding/json"
	"log"
	"os"
)

type User struct {
	Name    string
	Age     int
	Country string
}
type Users []User

func WriteToJson(filePath string, users Users) error {
	// TODO: answer here
	return nil
}

func main() {
	err := WriteToJson("./users.json", Users{
		{
			Name:    "levi",
			Age:     32,
			Country: "Indonesia",
		},
		{
			Name:    "jeff",
			Age:     30,
			Country: "USA",
		},
	})
	if err != nil {
		log.Fatal(err)
	}
}
