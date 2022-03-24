package main

import (
	"errors"
	"fmt"
	"os"
)

func IsFileExists(filename string) (bool, error) {
	file, err := os.Open(filename)
	defer file.Close()
	if err != nil {
		// Wrapping error
		return false, fmt.Errorf("Error in IsFileExists, err: %w", err)
	}

	return true, nil
}

// Ketika kita membungkus error, kita tidak dapat mengecek jenis error dengan menggunakan ==.
// Cara yang dapat digunakana adalah dengan menggunakan errors.Is

func main() {
	_, err := IsFileExists("file.txt")
	if err != nil {
		// Print error dari IsFileExists dengan informasi tambahan
		fmt.Println(err)

		// Mengecek jenis error
		if errors.Is(err, os.ErrNotExist) {
			fmt.Println("Error file tidak tersedia")
		}
	}
}
