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

// Struct untuk menyimpan data json
type user struct {
	Name   string `json:"name"`
	Age    int    `json:"age"`
	Email  string `json:"email"`
	Role   string `json:"role"`
	Status string `json:"status"`
}

// Format output agar lebih rapi dan mudah dibaca:
func (s user) String() string {
	return fmt.Sprintf("name: %v\nage: %v\nemail: %v\nrole: %v\nstatus: %v\n", s.Name, s.Age, s.Email, s.Role, s.Status)
}

func main() {
	fileName := "read"
	json, err := readJSON(fileName)
	if err != nil {
		fmt.Println(err)
	}

	for _, user := range json {
		fmt.Println(user)
	}
}

func readJSON(fileName string) ([]user, error) {
	path, err := filepath.Abs(fileName + ".json")
	if err != nil {
		return nil, err
	}
	file, err := openFile(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	users := []user{}

	byteValue, _ := ioutil.ReadAll(file)
	json.Unmarshal(byteValue, &users)
	return users, nil
}

func openFile(path string) (*os.File, error) {
	file, err := os.OpenFile(path, os.O_RDWR, os.ModePerm)
	if err != nil {
		return nil, err
	}
	return file, nil
}
