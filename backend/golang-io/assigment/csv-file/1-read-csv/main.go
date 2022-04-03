package main

import (
	"fmt"
	"log"
)

func main() {
	CSVRead()
	data, err := CSVReadAll()
	if err != nil {
		log.Fatal(err)
	}

	for _, record := range data {
		user := Person{
			FirstName: record[0],
			LastName:  record[1],
			Domicile:  record[2],
		}
		fmt.Println(user)
	}
}
