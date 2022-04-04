package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
)

//jalankan kode ini di dalam directory yang sama dengan lokasi kode
//hal ini perlu dilakukan agar file yang dicari dapat ditemukan

// TODO: answer here

func (s user) String() string {
	return fmt.Sprintf("name: %v\nschool: %v\nclass: %v\nphone: %v\nscore: %v\n", s.Name, s.School, s.Class, s.Phone, s.Score)
}

func main() {
	fileName := "data/read"
	json, err := readJSON(fileName)
	if err != nil {
		fmt.Println(err)
	}

	for _, user := range json {
		fmt.Println(user)
	}
}

func readJSON(fileName string) ([]user, error) {
	[]user{}, nil // TODO: replace this
}

func openFile(path string) (*os.File, error) {
	file, err := os.OpenFile(path, os.O_RDWR, os.ModePerm)
	if err != nil {
		return nil, err
	}
	return file, nil
}
