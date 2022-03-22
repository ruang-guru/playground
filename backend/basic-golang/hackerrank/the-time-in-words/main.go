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
    result := ""
    resultHour:= ""
    resultMinute:= ""
    inBetween :=""
    var toWord = map[int32]string{
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
      20: "twenty ",
    }

    if m < 30 {
        inBetween = "minute past"
        if m == 15 {
            resultMinute = "quarter"
             inBetween = " past"
        }else if m>20 {
            m -=20
            resultMinute = toWord[20]+toWord[m]
        }else {
            resultMinute = toWord[m]
        }
    } else if m>30 {
        inBetween = "minutes to"
        m = 60-m
        if m == 15 {
            resultMinute = "quarter"
        }else if m>20 {
            m -=20
            resultMinute = toWord[20]+" "+toWord[m]
        }else {
            resultMinute = toWord[m]
        }
        h++
    }

    if h>12 {
        h-=12
        resultHour = toWord[h]
    }else {
        resultHour = toWord[h]
    }

    if m ==0 {
        result = resultHour + " o' clock"
    }else{    
        result = resultMinute+" "+inBetween+" "+resultHour
    }

    return result // TODO: replace this
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
