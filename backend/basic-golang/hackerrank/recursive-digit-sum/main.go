package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

/*
 * Complete the 'superDigit' function below.
 *
 * The function is expected to return an INTEGER.
 * The function accepts following parameters:
 *  1. STRING n
 *  2. INTEGER k
 */

//Problem statement: https://www.hackerrank.com/challenges/recursive-digit-sum/problem
func superDigit(n string, k int32) int32 {
	//beginanswer
	sumDigit := func(n string) int64 {
		sum := int64(0)
		for _, c := range n {
			digit := int64(c) - '0'
			sum += digit
		}
		return sum
	}

	if len(n) == 1 && k == 1 {
		res, _ := strconv.Atoi(n)
		return int32(res)
	} else {
		return superDigit(fmt.Sprint(sumDigit(n)*int64(k)), 1)
	}
	//endanswer return 0
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)

	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 16*1024*1024)

	firstMultipleInput := strings.Split(strings.TrimSpace(readLine(reader)), " ")

	n := firstMultipleInput[0]

	kTemp, err := strconv.ParseInt(firstMultipleInput[1], 10, 64)
	checkError(err)
	k := int32(kTemp)

	result := superDigit(n, k)

	fmt.Fprintf(writer, "%d\n", result)

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
