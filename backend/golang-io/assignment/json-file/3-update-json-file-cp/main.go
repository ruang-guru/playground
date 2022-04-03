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

func main() {
	fileName := "siswa"
	arif := student{Name: "arif", Score: 90, Class: "c"}
	andi := student{Name: "andi", Score: 85, Class: "c"}
	newData := []student{arif, andi}
	json, _ := updateJSON(fileName, newData)
	for _, student := range json {
		fmt.Println(student)
	}

	//comment baseData and reset to show change in json file
	baseData := []student{
		{
			Name:  "aldo",
			Score: 90,
			Class: "C",
		},
		{
			Name:  "aldi",
			Score: 94,
			Class: "C",
		},
		{
			Name:  "ado",
			Score: 93,
			Class: "C",
		},
	}
	reset(fileName, baseData)
}

func updateJSON(fileName string, newData []student) ([]student, error) {
	return []student{}, nil // TODO: replace this
}

func openFile(path string) (*os.File, error) {
	file, err := os.OpenFile(path, os.O_RDWR|os.O_CREATE, os.ModePerm)
	if err != nil {
		return nil, err
	}
	return file, nil
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

	student := []student{}

	byteValue, _ := ioutil.ReadAll(file)
	json.Unmarshal(byteValue, &student)
	return student, nil
}

func (s student) String() string {
	return fmt.Sprintf("name: %v\nscore: %v\nclass: %v\n", s.Name, s.Class, s.Score)
}

func reset(fileName string, baseData []student) {
	path := fileName + ".json"
	os.Remove(path)
	updateJSON(fileName, baseData)
}
