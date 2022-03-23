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
 * Complete the 'birthdayCakeCandles' function below.
 *
 * The function is expected to return an INTEGER.
 * The function accepts INTEGER_ARRAY candles as parameter.
 */

func birthdayCakeCandles(candles []int32) int32 {
    // Write your code here
	var result int32
	var a int32
	a =candles[0]
	
for i := 0; i < len(candles); i++ {
	if a<candles[i] {
		a = candles[i]
	}
}
for i := 0; i < len(candles); i++ {
	if a==candles[i] {
		result++
	}
}

return result
}

func main() {
    reader := bufio.NewReaderSize(os.Stdin, 16 * 1024 * 1024)

    stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
    checkError(err)

    defer stdout.Close()

    writer := bufio.NewWriterSize(stdout, 16 * 1024 * 1024)

    candlesCount, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
    checkError(err)

    candlesTemp := strings.Split(strings.TrimSpace(readLine(reader)), " ")

    var candles []int32

    for i := 0; i < int(candlesCount); i++ {
        candlesItemTemp, err := strconv.ParseInt(candlesTemp[i], 10, 64)
        checkError(err)
        candlesItem := int32(candlesItemTemp)
        candles = append(candles, candlesItem)
    }

    result := birthdayCakeCandles(candles)

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
