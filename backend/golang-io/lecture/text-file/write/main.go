package main

// importing the packages
import (
	"fmt"
	"log"
	"os"
)

func CreateFile() {

	//membuat file
	file, err := os.Create("write.txt")

	if err != nil {
		log.Fatalf("failed creating file: %s", err)
	}

	//menutup file sebelum fungsi selesai dijalankan
	defer file.Close()

	len, err := file.WriteString("Write anything here\n" +
		"This program demonstrates reading and writing\n" +
		"operations to a file in Go")

	if err != nil {
		log.Fatalf("failed writing to file: %s", err)
	}

	fmt.Printf("File Name: %s\n", file.Name())
	fmt.Printf("Length: %d bytes\n", len)
}

func main() {
	CreateFile()
	//reset() //uncoment this to remove the created text file
}

// reference : https://www.geeksforgeeks.org/how-to-read-and-write-the-files-in-golang/
// textfile yang dibuat akan tampil dalam folder ini

func reset() {
	os.Remove("write.txt")
}
