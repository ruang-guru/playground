package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

/*
 * Complete the 'isValid' function below.
 *
 * The function is expected to return a STRING.
 * The function accepts STRING s as parameter.
 */

//Problem statement: https://www.hackerrank.com/challenges/sherlock-and-valid-string/problem
func isValid(s string) string {
	// Write your code here
	//beginanswer
	count := make(map[rune]int)
	for _, c := range s {
		count[c]++
	}

	uniqCounts := make(map[int]int)
	for _, v := range count {
		uniqCounts[v] += 1
	}

	if len(uniqCounts) == 1 {
		return "YES"
	} else if len(uniqCounts) == 2 {
		keys := []int{}
		values := []int{}
		for k, v := range uniqCounts {
			keys = append(keys, k)
			values = append(values, v)
		}

		for i := 0; i < 2; i++ {
			if keys[i] == 1 && values[i] == 1 {
				//since the count is 1, we can just remove it
				return "YES"
			}
			//need to reduce by one
			if keys[i]-keys[1-i] == 1 {
				//only one item can be reduced
				if values[i] == 1 {
					return "YES"
				}
			}
		}
	}
	return "NO"
	//endanswer return ""
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)

	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 16*1024*1024)

	s := readLine(reader)

	result := isValid(s)

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
