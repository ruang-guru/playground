package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
)

// Ada banyak cara untuk membaca text file line by line.

func main() {
	// ReadWithReadLine()
	ReadWithScanner()
}

// Dengan bufio.Reader.ReadLine
func ReadWithReadLine() {
	file, err := os.Open("main.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	reader := bufio.NewReader(file)
	for {
		line, _, err := reader.ReadLine()
		if err == io.EOF {
			break
		}
		fmt.Printf("%s \n", line)
	}
}

// Dengan bufio.Scanner
func ReadWithScanner() {
	file, err := os.Open("main.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
}
