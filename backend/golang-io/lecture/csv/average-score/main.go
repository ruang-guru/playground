package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"path/filepath"
	"strconv"
)

//jalankan dari dalam folder average-score

//membuka file
func openFile(csvName string) (*os.File, error) {
	path, err := filepath.Abs("../data/" + csvName + ".csv")

	studentsFile, err := os.OpenFile(path, os.O_RDWR|os.O_CREATE, os.ModePerm)
	if err != nil {
		return nil, err
	}

	return studentsFile, nil
}

//mendapatkan nilai dari csv
func loadCSV(dbName string) ([][]string, error) {
	f, err := openFile(dbName)
	if err != nil {
		return nil, err
	}

	defer f.Close()

	csvReader := csv.NewReader(f)
	records, err := csvReader.ReadAll()
	if err != nil {
		return nil, err
	}
	return records, nil
}

func main() {
	fmt.Println(getClassAverageScore("B"))
}

//mendapatkan nilai
func getScores(students [][]string) []int {
	scores := []int{}
	for i := 1; i < len(students); i++ {
		scoreStr, _ := strconv.Atoi(students[i][1])
		scores = append(scores, scoreStr)

	}
	return scores
}

func getClassAverageScore(class string) float32 {
	dbName := "siswa" + class
	data, err := loadCSV(dbName)
	if err != nil {
		fmt.Println("error : ", err)
	}
	average := getAverageScore(getScores(data))
	return average
}

func getAverageScore(scores []int) float32 {
	sum := 0
	for _, score := range scores {
		sum += score
	}
	fmt.Println(sum)
	fmt.Println(len(scores))
	averageScore := float32(sum / len(scores))
	return averageScore
}
