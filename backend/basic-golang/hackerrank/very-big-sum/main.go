
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
 * Complete the 'aVeryBigSum' function below.
 *
 * The function is expected to return a LONG_INTEGER.
 * The function accepts LONG_INTEGER_ARRAY ar as parameter.
 */

func aVeryBigSum(ar []int64) int64 {
    // Write your code here
	var r int64
	for i := 0; i < len(ar); i++ {
	r +=ar[i]	
	}

	
return r	
};

func main() {
    reader := bufio.NewReaderSize(os.Stdin, 16 * 1024 * 1024)

    stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
    checkError(err)

    defer stdout.Close()

    writer := bufio.NewWriterSize(stdout, 16 * 1024 * 1024)

    arCount, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
    checkError(err)

    arTemp := strings.Split(strings.TrimSpace(readLine(reader)), " ")

    var ar []int64

    for i := 0; i < int(arCount); i++ {
        arItem, err := strconv.ParseInt(arTemp[i], 10, 64)
        checkError(err)
        ar = append(ar, arItem)
    }

    result := aVeryBigSum(ar)

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
