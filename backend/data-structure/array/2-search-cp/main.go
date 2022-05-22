package main

import (
	"errors"
	"fmt"
)

func main() {
	var cars1 = []string{"Toyota", "Honda", "Nissan", "BMW", "Chevy", "Ford"}
	var cars2 = []string{"Ford", "BMW", "Audi", "Mercedes"}

	result, err := SearchMatch(cars1, cars2)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Matched car brand: ", result)
	}
}

func SearchMatch(arr1 []string, arr2 []string) ([]string, error) {
	matchedString := make([]string, 0)
	for _, val1 := range arr1 {
		for _, val2 := range arr2 {
			if val1 == val2 {
				matchedString = append(matchedString, val1)
			}
		}
	}

	if len(matchedString) == 0 {
		return []string{}, errors.New("no match")
	}
	return matchedString, nil // TODO: replace this
}
