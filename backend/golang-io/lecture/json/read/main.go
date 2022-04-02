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

type student struct {
	Name  string `json:"nama"`
	Score int    `json:"nilai"`
	Class string `json:"kelas"`
}

func (s student) String() string {
	return fmt.Sprintf("name: %v\nscore: %v\nclass: %v\n", s.Name, s.Class, s.Score)
}

func main() {
	fileName := "siswa"
	json, _ := readJSON(fileName)
	for _, student := range json {
		fmt.Println(student)
	}
}

func readJSON(fileName string) ([]student, error) {
	path, err := filepath.Abs(fileName + ".json")
	if err != nil {
		return nil, err
	}
	file, err := openFile(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	students := []student{}

	byteValue, _ := ioutil.ReadAll(file)
	json.Unmarshal(byteValue, &students)
	return students, nil
}

func openFile(path string) (*os.File, error) {
	file, err := os.OpenFile(path, os.O_RDWR, os.ModePerm)
	if err != nil {
		return nil, err
	}
	return file, nil
}
