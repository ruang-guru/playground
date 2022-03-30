package main

import (
	"errors"
	"fmt"
)

// Dari contoh yang telah diberikan, coba kamu gunakan errors.Is untuk mengecek jenis error tertentu.
// Modifikasikan pada kode dari wrapping error
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
		// Cek apakah err merupakan jenis error ErrDataNotFound
		// TODO: answer here
	}
}
