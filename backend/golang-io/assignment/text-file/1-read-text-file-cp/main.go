package main

import (
	"fmt"
	"io/ioutil"
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
	data, err := ioutil.ReadFile(name)
	if err != nil {
		fmt.Println(err)
	}
	return FileData{
		Name: name,
		Size: len(data),
		Data: data,
	}, err // TODO: replace this
}
