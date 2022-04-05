package main

import (
	"fmt"
	"io/ioutil"
	"log"
)

func main() {
	fmt.Println("dummyCommit")
}

// Gunakan struct untuk menyimpan data file nya ya
type FileData struct {
	Name string
	Size int
	Data []byte
}

func ReadFile(name string) (FileData, error) {
	// nama text file yang ingin dibaca
	fileName := "./read.txt"

	//membaca text file
	data, err := ioutil.ReadFile(fileName)
	if err != nil {
		log.Panicf("failed reading data from file: %s", err)
	}

	fileData := FileData{
		Name: fileName,
		Size: len(data),
		Data: data,
	}
	return fileData, nil
}
