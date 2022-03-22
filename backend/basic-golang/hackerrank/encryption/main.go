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
	//beginanswer
	root := math.Sqrt(float64(len(s)))
	width := int(math.Ceil(root))
	height := int(math.Floor(root))
	if width*height < len(s) {
		height++
	}

	m := make([][]rune, height)
	for i := range m {
		m[i] = make([]rune, width)
	}

	x := 0
	y := 0
	for _, c := range s {
		m[y][x] = c
		if x == width-1 {
			x = 0
			y++
		} else {
			x++
		}
	}

	var result strings.Builder
	for i := 0; i < width; i++ {
		for j := 0; j < height; j++ {
			if m[j][i] != 0 {
				result.WriteRune(m[j][i])
			}
		}
		if i != width-1 {
			result.WriteString(" ")
		}
	}

	return result.String()
	//endanswer return ""
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
