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
 * Complete the 'timeInWords' function below.
 *
 * The function is expected to return a STRING.
 * The function accepts following parameters:
 *  1. INTEGER h
 *  2. INTEGER m
 */

//Problem statement: https://www.hackerrank.com/challenges/the-time-in-words/problem

func timeInWords(h int32, m int32) string {
	// Write your code here
	//beginanswer
	numberToWord := map[int32]string{
		1:  "one",
		2:  "two",
		3:  "three",
		4:  "four",
		5:  "five",
		6:  "six",
		7:  "seven",
		8:  "eight",
		9:  "nine",
		10: "ten",
		11: "eleven",
		12: "twelve",
		13: "thirteen",
		14: "fourteen",
		15: "fifteen",
		16: "sixteen",
		17: "seventeen",
		18: "eighteen",
		19: "nineteen",
		20: "twenty",
		30: "thirty",
		40: "forty",
		50: "fifty",
		60: "sixty",
	}
	for i := int32(1); i <= 60; i++ {
		if _, exist := numberToWord[i]; !exist {
			numberToWord[i] = fmt.Sprintf("%s %s", numberToWord[i-i%10], numberToWord[i%10])
		}
	}
	numberToWord[15] = "quarter"
	numberToWord[30] = "half"

	if m == 0 {
		return fmt.Sprintf("%s o' clock", numberToWord[h])
	}
	if m <= 30 {
		if m == 15 || m == 30 {
			return fmt.Sprintf("%s past %s", numberToWord[m], numberToWord[h])
		} else if m == 1 {
			return fmt.Sprintf("%s minute past %s", numberToWord[m], numberToWord[h])
		} else {
			return fmt.Sprintf("%s minutes past %s", numberToWord[m], numberToWord[h])
		}
	} else {
		if m == 45 {
			return fmt.Sprintf("quarter to %s", numberToWord[h+1])
		} else if m == 1 {
			return fmt.Sprintf("%s minute to %s", numberToWord[60-m], numberToWord[h+1])
		} else {
			return fmt.Sprintf("%s minutes to %s", numberToWord[60-m], numberToWord[h+1])
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

	hTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
	checkError(err)
	h := int32(hTemp)

	mTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
	checkError(err)
	m := int32(mTemp)

	result := timeInWords(h, m)

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
