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
	file, err := os.OpenFile(fileName, os.O_RDONLY, 0644)
	if err != nil {
		fmt.Println(err)
	}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		*arr = append(*arr, scanner.Text())
	}
	return err // TODO: replace this
}

func ScanToMap(dataMap map[string]string, fileName string) error {
	file, err := os.OpenFile(fileName, os.O_RDONLY, 0644)
	if err != nil {
		fmt.Println(err)
	}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		data := strings.Split(scanner.Text(), ",")
		dataMap[data[0]] = data[1]
	}
	return err // TODO: replace this
}
