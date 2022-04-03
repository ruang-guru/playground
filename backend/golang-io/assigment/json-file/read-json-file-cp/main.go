package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
)

type User struct {
	Name    string `json:"name"`
	Age     int    `json:"age"`
	Country string `json:"country"`
}
type Users []User

func ReadUsers(filePath string) (Users, error) {
	panic("Not yet implemented") // TODO: answer here
}

func main() {
	users, err := ReadUsers("./users.json")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%+v\n", users)
}
