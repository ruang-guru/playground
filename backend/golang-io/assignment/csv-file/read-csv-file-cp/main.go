package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strconv"
)

type User struct {
	Name    string
	Age     int
	Country string
}
type Users []User

func ReadUsers(filePath string) (Users, error) {
	// TODO: answer here
}

func main() {
	users, err := ReadUsers("./users.csv")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%+v\n", users)
}
