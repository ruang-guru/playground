package main

import (
	"errors"
	"fmt"
)

// Dari contoh yang telah diberikan, kamu dapat mencoba untuk membungkus error ke dalam error yang lain.
// Kamu dapat mencoba untuk membungkus error dari GetAge pada function IsEligibleToVaccine.
// Kamu dapat memberikan informasi tambahan pada pembungkusan error tersebut, misalnya nama function.

var ErrDataNotFound = errors.New("error data not found")

func GetAge(data map[string]int, name string) (int, error) {
	if _, ok := data[name]; !ok {
		return 0, ErrDataNotFound
	}

	return data[name], nil
}

func IsEligibleToVaccine(data map[string]int, name string) (bool, error) {
	age, err := GetAge(data, name)
	if err != nil {
		// TODO: answer here
	}
	if age < 15 {
		return false, nil
	}

	return true, nil
}

func main() {
	data := map[string]int{
		"Roger":  60,
		"Kamala": 5,
		"Peter":  20,
	}
	_, err := IsEligibleToVaccine(data, "Tony")
	if err != nil {
		// Coba kalian print pesan error dari error IsEligibleToVaccine dan print pesan error yang dibungkus
		// TODO: answer here
	}
}
