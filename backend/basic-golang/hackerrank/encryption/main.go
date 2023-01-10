package main

import (
	"bufio"
	"fmt"
	"io"
	"math"
	"math/rand"
	"os"
	"strings"
)

/*
 * Complete the 'encryption' function below.
 *
 * The function is expected to return a STRING.
 * The function accepts STRING s as parameter.
 */

func encryption(s string) string {
    // Write your code here

	rows := math.Floor(math.Sqrt(float64(len(s))))
	collumn := math.Ceil(math.Sqrt(float64(len(s))))
	var slice []string
	
	for i := 0; i < len(s); i++ {
		slice =append(slice,string(s[i]))
	}

	for{
		for i := 0; i < int(rows); i++ {
		temp:= rand.Intn(int(collumn))	
		fmt.Print(s[temp])
		
		}

		count:= 1
		if count ==int(collumn) {
			continue
		}
		if len(slice)==0 {
			break
		}
	}

}

func main() {
    reader := bufio.NewReaderSize(os.Stdin, 16 * 1024 * 1024)

    stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
    checkError(err)

    defer stdout.Close()

    writer := bufio.NewWriterSize(stdout, 16 * 1024 * 1024)

    s := readLine(reader)

    result := encryption(s)

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
