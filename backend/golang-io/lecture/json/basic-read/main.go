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

type user struct {
	Name string
	Age  int
}

type student struct {
	Name  string
	Score int
	Class string
}

func openFile(path string) (*os.File, error) {
	file, err := os.OpenFile(path, os.O_RDWR|os.O_CREATE, os.ModePerm)
	if err != nil {
		return nil, err
	}
	return file, nil
}

func main() {
	fileName := "user"
	json, _ := readJSON(fileName)
	fmt.Println(json)
}

func readJSON(fileName string) (user, error) {
	path, err := filepath.Abs(fileName + ".json")
	if err != nil {
		return user{}, err
	}
	file, err := openFile(path)
	if err != nil {
		return user{}, err
	}
	defer file.Close()

	student := user{}

	byteValue, _ := ioutil.ReadAll(file)
	json.Unmarshal(byteValue, &student)

	return student, nil
}
