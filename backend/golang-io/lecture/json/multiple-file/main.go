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

type classesAverageScore struct {
	Name         string  `json:"kelas"`
	AverageScore float32 `json:"rata-rata"`
	TotalStudent int     `json:"total-siswa"`
}

func openFile(path string) (*os.File, error) {
	file, err := os.OpenFile(path, os.O_RDWR|os.O_CREATE, os.ModePerm)
	if err != nil {
		return nil, err
	}
	return file, nil
}

func main() {
	classesName := []string{"B", "C"}
	classes := getClassesData(classesName)
	for _, class := range classes {
		fmt.Println(class)
	}
	writeJSON(classes)
}

func getClassesData(classesName []string) []classesAverageScore {
	classesData := []classesAverageScore{}
	for _, className := range classesName {
		students, _ := getStudent(className)
		sum := 0
		for _, student := range students {
			sum += student.Score
		}
		average := float32(sum / len(students))
		classData := classesAverageScore{
			Name:         className,
			AverageScore: average,
			TotalStudent: len(students),
		}
		classesData = append(classesData, classData)
	}
	return classesData
}

//membaca data student yang ada dalam file json
func getStudent(class string) ([]student, error) {
	path, err := filepath.Abs("siswa" + class + ".json")
	if err != nil {
		return nil, err
	}
	file, err := openFile(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	student := []student{}

	byteValue, err := ioutil.ReadAll(file)
	if err != nil {
		return nil, err
	}
	json.Unmarshal(byteValue, &student)
	return student, nil
}

func writeJSON(classes []classesAverageScore) error {
	path, err := filepath.Abs("Classes average score " + ".json")
	if err != nil {
		return err
	}
	file, err := openFile(path)
	if err != nil {
		return err
	}
	defer file.Close()
	jsonData, _ := json.Marshal(classes)
	ioutil.WriteFile(path, jsonData, 0644)
	return nil
}

func (c classesAverageScore) String() string {
	return fmt.Sprintf("class: %v\naverageScore: %v\ntotalStudent: %v\n", c.Name, c.AverageScore, c.TotalStudent)
}
