package main

import (
	"encoding/base64"
	"fmt"
)

// print encode string 'tokenrahasia' kedalam base64
func main() {
	str := "tokenrahasia"

	encodedStr := base64.StdEncoding.EncodeToString([]byte(str))

	fmt.Println(encodedStr)
	fmt.Println()
}
