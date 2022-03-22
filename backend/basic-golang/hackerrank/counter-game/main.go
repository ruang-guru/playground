package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

//beginanswer
//not optimal O(n) search, but enough for now
func contains(arr []int64, n int64) bool {
	for _, v := range arr {
		if v == n {
			return true
		}
	}
	return false
}

//endanswer nop

/*
 * Complete the 'counterGame' function below.
 *
 * The function is expected to return a STRING.
 * The function accepts LONG_INTEGER n as parameter.
 */

//Problem statement: https://www.hackerrank.com/challenges/counter-game
func counterGame(n int64) string {
	// Write your code here
	//beginanswer
	powerOfTwos := make([]int64, 0)
	for i := int64(1); i <= n; i *= 2 {
		powerOfTwos = append(powerOfTwos, i)
	}

	turn := 0
	for n > 1 {
		turn += 1
		if contains(powerOfTwos, n) {
			n /= 2
			continue
		}

		prevPowerOfTwo := int64(1)
		for i := len(powerOfTwos) - 1; i >= 0; i-- {
			if n > powerOfTwos[i] {
				prevPowerOfTwo = powerOfTwos[i]
				break
			}
		}
		n -= prevPowerOfTwo
	}

	if turn == 0 {
		return "Louise"
	} else {
		if turn%2 == 1 {
			return "Louise"
		} else {
			return "Richard"
		}
	}

	//endanswer return ""
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)

	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 16*1024*1024)

	tTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
	checkError(err)
	t := int32(tTemp)

	for tItr := 0; tItr < int(t); tItr++ {
		n, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
		checkError(err)

		result := counterGame(n)

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
