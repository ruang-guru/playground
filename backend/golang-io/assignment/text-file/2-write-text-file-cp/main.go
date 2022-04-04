package main

import (
	"log"
	"os"
)

// dalam test ini terdapat fungsi os.Remove ya. itu automatis nge remove file yang telah dibuat
// Untuk keperluan testing
func WriteFile(fileName string, fileData string) error {
	file, err := os.Create(fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	_, err = file.WriteString(fileData)
	if err != nil {
		log.Fatal(err)
	}

	return err // TODO: replace this
}
