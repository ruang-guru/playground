package main

// importing the packages
import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"
)

func UpdateFile(newWord string) {
	// nama text file yang ingin dibaca
	fileName := "read.txt"

	//membaca text file
	data, err := ioutil.ReadFile(fileName)
	if err != nil {
		log.Panicf("failed reading data from file: %s", err)
	}
	fmt.Println("Before update :")
	fmt.Printf("File Name: %s", fileName)
	fmt.Printf("\nSize: %d bytes", len(data))
	fmt.Printf("\nData: %s\n\n", data)

	//memisahkan string berdasarkan spasi atau " "
	lines := strings.Split(string(data), " ")

	for i, line := range lines {
		//mengecek apakah dalam baris ada string "read"
		if strings.Contains(line, "read") {
			lines[i] = newWord
		}
	}
	output := strings.Join(lines, " ")
	err = ioutil.WriteFile(fileName, []byte(output), 0644)
	if err != nil {
		log.Fatalln(err)
	}
	data, err = ioutil.ReadFile(fileName)
	if err != nil {
		log.Panicf("failed reading data from file: %s", err)
	}
	fmt.Println("After update :")
	fmt.Printf("File Name: %s", fileName)
	fmt.Printf("\nSize: %d bytes", len(data))
	fmt.Printf("\nData: %s\n", data)

}

func main() {
	word := "update"
	UpdateFile(word)
	reset(word) // delete this to show change in txt file
}

//reference : https://www.geeksforgeeks.org/how-to-read-and-write-the-files-in-golang/

//reset textfile
func reset(newWord string) {
	// nama text file yang ingin dibaca
	fileName := "read.txt"

	//membaca text file
	data, err := ioutil.ReadFile(fileName)
	if err != nil {
		log.Panicf("failed reading data from file: %s", err)
	}

	//memisahkan string berdasarkan spasi atau " "
	lines := strings.Split(string(data), " ")

	for i, line := range lines {
		//mengecek apakah dalam baris ada string "read"
		if strings.Contains(line, newWord) {
			lines[i] = "read"
		}
	}
	output := strings.Join(lines, " ")
	err = ioutil.WriteFile(fileName, []byte(output), 0644)
	if err != nil {
		log.Panicf("failed reading data from file: %s", err)
	}
}
