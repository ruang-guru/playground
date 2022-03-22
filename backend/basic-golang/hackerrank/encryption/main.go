package main

import (
	"bufio"
	"fmt"
	"io"
	"math"
	"os"
	"strings"
)

/*
 * Complete the 'encryption' function below.
 *
 * The function is expected to return a STRING.
 * The function accepts STRING s as parameter.
 */

//Problem statement: https://www.hackerrank.com/challenges/encryption/problem

func encryption(s string) string {
	// Write your code here
	s = strings.Replace(s, " ", "", -1)
	result := ""
	length := len(s)
	root := math.Sqrt(float64(length))
	var row, col int
	if root == math.Floor(root) {
		row = int(root)
		col = int(root)
	} else {
		row = int(math.Floor(root))
		col = row + 1

		if row*col < length {
			row += 1
		}
	}
	for i := 0; i <= col; i++ {
		for j := 0; j <= row; j++ {
			ind := i + j*col
			if ind < length {
				result += s[ind : ind+1]
			}
		}
		result += " "
	}
	return result // TODO: replace this
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)

	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 16*1024*1024)

	s := readLine(reader)

	result := encryption(s)

	fmt.Fprintf(writer, "%s\n", result)

	writer.Flush()
}

func readLine(reader *bufio.Reader) string {
	str, _, err := reader.ReadLine()
	if err == io.EOF {
		return ""
	}

	return strings.TrimRight(string(str), "\r\n")
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
