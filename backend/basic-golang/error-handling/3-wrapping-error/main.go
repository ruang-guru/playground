package main

import (
	"errors"
	"fmt"
	"os"
)

// Untuk memberikan informasi atau konteks tambahan pada error, kita dapat membungkus error ke dalam error yang lain.
// Sebagai contoh adalah pada function IsFileExists terdapat pemanggilan method os.Open.
// Jika os.Open mengembalikan nilai error, maka pada function IsFileExists akan mengembalikan nilai false dan error.
// Untuk memberikan informasi tambahan pada error yang terjadi di IsFileExists, kita dapat membungkus error dari os.Open ke dalam error yang
// dikembalikan oleh IsFileExists
func IsFileExists(filename string) (bool, error) {
	file, err := os.Open(filename)
	defer file.Close()
	if err != nil {
		// Wrapping error
		return false, fmt.Errorf("Error in IsFileExists, err: %w", err)
	}

	return true, nil
}

func main() {
	_, err := IsFileExists("file.txt")
	if err != nil {
		// Print error dari IsFileExists dengan informasi tambahan
		fmt.Println(err)

		// Untuk mendapatkan error yang dibungkus, kita dapat menggunakan function Unwrap.
		if errWrap := errors.Unwrap(err); errWrap != nil {
			// Print error dari os.Open()
			fmt.Println(errWrap)
		}
	}
}
