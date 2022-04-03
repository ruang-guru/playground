package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Print("Hello World")
}

func ScanToArray(arr *[]string, fileName string) error {
	file, err := os.Open(fileName)
	if err != nil {
		return err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		*arr = append(*arr, scanner.Text())
	}
	return nil
}

func ScanToMap(dataMap map[string]string, fileName string) error {
	file, err := os.Open(fileName)
	if err != nil {
		return err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	// scanner.Split(bufio.ScanWords)
	for scanner.Scan() {
		text := strings.Split((scanner.Text()), ",")

		dataMap[text[0]] = text[1]
	}
	return nil
}
