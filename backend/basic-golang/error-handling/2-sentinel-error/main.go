package main

import (
	"errors"
	"fmt"
)

// Sentinel error merupakan error yang dibungkus dalam suatu variabel.
// Untuk membuat sentinel error kita dapat menggunakan errors.New, fmt.Errorf, maupun dari custom error.
var ErrDataNotFound = errors.New("error data not found")

func GetAge(data map[string]int, name string) (int, error) {
	if _, ok := data[name]; !ok {
		return 0, ErrDataNotFound
	}

	return data[name], nil
}

func main() {
	peopleAge := map[string]int{
		"John": 20,
		"Raz":  8,
		"Tony": 40,
	}

	_, err := GetAge(peopleAge, "Roger")
	if err != nil {
		// Sentinel error memberikan kita keleluasaan dalam memberikan response terhadap error tersebut.
		if err == ErrDataNotFound {
			fmt.Printf("Data %s is not found, error: %s\n", "Roger", err.Error())
		}
	}

}
