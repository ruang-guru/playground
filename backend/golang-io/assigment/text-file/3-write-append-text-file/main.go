package main

import (
	"io/ioutil"
	"log"
	"os"
)

func main() {
	// Silahkan kalian cari tau arti dari os.O_APPEND dan os.O_WRONLY dan 0644 ini apa
	f, err := os.OpenFile("namaFile.txt", os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		log.Printf("failed opening file: %s", err)
	}

	defer f.Close()
	if _, err = f.WriteString("Badabing bada bdong "); err != nil {
		log.Printf("failed writing to file: %s", err)
	}

	// read file again
	data, err := ioutil.ReadFile("namaFile.txt")
	if err != nil {
		log.Printf("failed reading file: %s", err)
	}
	log.Printf("%s", data)
}
