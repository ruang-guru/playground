package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"
)

func main() {
	fmt.Print("dummy commit")
}

func CSVToMap(data map[string]string, fileName string) (map[string]string, error) {
	// TODO: answer here
	file, err := os.Open(fileName)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	defer file.Close()
	reader := csv.NewReader(file)
	for {
		dataReader, err := reader.Read()
		if err == io.EOF {
			break
		}
		data[dataReader[0]] = dataReader[1]
	}

	return data, nil
}
