package main

import (
	"bufio"
	"fmt"
	"io"
	"math"
	"os"
	"strconv"
	"strings"
)

/*
 * Complete the 'counterGame' function below.
 *
 * The function is expected to return a STRING.
 * The function accepts LONG_INTEGER n as parameter.
 */

//Problem statement: https://www.hackerrank.com/challenges/counter-game
func counterGame(n int64) string {
	// Write your code here
	isLouise := false
	winner := ""
	if n == 1 {
		winner = "Richard"
	} else {
		pow := 10
		for n > 1 {
			powResBefore := int64(math.Pow(2, float64(pow)))
			for {
				checkNumber := int64(math.Pow(2, float64(pow)))
				if checkNumber > n {
					if powResBefore < n {
						n -= powResBefore
						isLouise = !isLouise
						pow = pow / 2
						break
					} else {
						powResBefore = checkNumber
						pow--
					}
				} else {
					if checkNumber == n {
						n /= 2
						isLouise = !isLouise
						pow = pow / 2
						break
					} else {
						powResBefore = checkNumber
						pow++
					}
				}
			}
		}
		if isLouise {
			winner = "Louise"
		} else {
			winner = "Richard"
		}
	}
	return winner // TODO: replace this
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
