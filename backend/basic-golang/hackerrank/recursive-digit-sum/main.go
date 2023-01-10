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

func superDigit(n string, k int32) int32 {
    // Write your code here
	
str := ""
    
for i := 0; i < int(k); i++ {
	str +=n
}

var slice[]int32
for i := 0; i < len(str); i++ {
	j,e:=strconv.Atoi(string(str[i]))
	if e==nil {
	fmt.Println("Data tidak valid")
	}
	slice = append(slice,int32(j))
}
temp:= int32(0)
for _, v := range slice {
	temp +=v
}


if temp>=10 {
	s:=strconv.Itoa(int(temp))
	superDigit(s,1)
} else {
	return temp
}


}

func main() {
    reader := bufio.NewReaderSize(os.Stdin, 16 * 1024 * 1024)

    stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
    checkError(err)

    defer stdout.Close()

    writer := bufio.NewWriterSize(stdout, 16 * 1024 * 1024)

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
