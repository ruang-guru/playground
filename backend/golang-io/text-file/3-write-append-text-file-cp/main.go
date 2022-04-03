package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Print("dummyCommit")
}

func AddString(fileName string, stringToAdd string) error {
	f, err := os.OpenFile(fileName, os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}

	defer f.Close()
	if _, err = f.WriteString(stringToAdd); err != nil {
		return err
	}
	return nil
}
