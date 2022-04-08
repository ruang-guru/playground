package main

// importing the packages
import (
	"fmt"
	"io/ioutil"
	"log"
)

func ReadFile() {
	// nama text file yang ingin dibaca
	fileName := "read.txt"

	//membaca text file
	data, err := ioutil.ReadFile(fileName)
	if err != nil {
		log.Panicf("failed reading data from file: %s", err)
	}
	fmt.Printf("File Name: %s", fileName)
	fmt.Printf("\nSize: %d bytes", len(data))
	fmt.Printf("\nData: %s\n", data)
}

func main() {
	ReadFile()
}

//reference : https://www.geeksforgeeks.org/how-to-read-and-write-the-files-in-golang/
