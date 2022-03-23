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
 * Complete the 'kangaroo' function below.
 *
 * The function is expected to return a STRING.
 * The function accepts following parameters:
 *  1. INTEGER x1
 *  2. INTEGER v1
 *  3. INTEGER x2
 *  4. INTEGER v2
 */

func kangaroo(x1 int32, v1 int32, x2 int32, v2 int32) string {
    // Write your code here
var k1[]int32
var k2[]int32
t1 :=v1+x1
t2 :=v2+x2
result := "NO"
	for i := 0; i < 1000; i++ {
		k1 = append(k1,t1)
		k2 = append(k2, t2)
		t1 +=v1
		t2 +=v2
	}

	for i := 0; i < 4; i++ {
		if k1[i]==k2[i] {
			result = "YES"
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

    firstMultipleInput := strings.Split(strings.TrimSpace(readLine(reader)), " ")

    x1Temp, err := strconv.ParseInt(firstMultipleInput[0], 10, 64)
    checkError(err)
    x1 := int32(x1Temp)

    v1Temp, err := strconv.ParseInt(firstMultipleInput[1], 10, 64)
    checkError(err)
    v1 := int32(v1Temp)

    x2Temp, err := strconv.ParseInt(firstMultipleInput[2], 10, 64)
    checkError(err)
    x2 := int32(x2Temp)

    v2Temp, err := strconv.ParseInt(firstMultipleInput[3], 10, 64)
    checkError(err)
    v2 := int32(v2Temp)

    result := kangaroo(x1, v1, x2, v2)

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
