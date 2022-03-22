package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"sort"
	"strconv"
	"strings"
)

/*
 * Complete the 'gamingArray' function below.
 *
 * The function is expected to return a STRING.
 * The function accepts INTEGER_ARRAY arr as parameter.
 */

//Problem statement: https://www.hackerrank.com/challenges/an-interesting-game-1/problem
func gamingArray(arr []int) string {
	// Write your code here
	// beginanswer
	positions := make(map[int]int)
	for i := 0; i < len(arr); i++ {
		positions[arr[i]] = i
	}
	sorted := make([]int, len(arr))
	copy(sorted, arr)
	sort.Slice(sorted, func(i, j int) bool {
		return sorted[i] > sorted[j]
	})

	turn := 0
	for len(arr) > 0 {
		turn += 1

		max := -1
		for {
			max = sorted[0]
			sorted = sorted[1:]
			if positions[max] < len(arr) {
				break
			}
		}

		deleteFrom := positions[max]
		arr = arr[:deleteFrom]
	}
	if turn%2 == 0 {
		return "ANDY"
	}
	return "BOB"
	//endanswer return ""
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)

	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 16*1024*1024)

	gTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
	checkError(err)
	g := int32(gTemp)

	for gItr := 0; gItr < int(g); gItr++ {
		arrCount, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
		checkError(err)

		arrTemp := strings.Split(strings.TrimSpace(readLine(reader)), " ")

		var arr []int

		for i := 0; i < int(arrCount); i++ {
			arrItemTemp, err := strconv.ParseInt(arrTemp[i], 10, 64)
			checkError(err)
			arrItem := int(arrItemTemp)
			arr = append(arr, arrItem)
		}

		result := gamingArray(arr)

		fmt.Fprintf(writer, "%s\n", result)
	}

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
