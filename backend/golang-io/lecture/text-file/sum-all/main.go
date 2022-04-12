package main

// importing the packages
import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"strings"
)

func ReadFile() []byte {
	// nama text file yang ingin dibaca
	fileName := "input.txt" //nilai dipisahkan oleh spasi

	//membaca text file
	data, err := ioutil.ReadFile(fileName)
	if err != nil {
		log.Panicf("failed reading data from file: %s", err)
	}
	fmt.Printf("File Name: %s", fileName)
	fmt.Printf("\nSize: %d bytes", len(data))
	fmt.Printf("\nData: %s\n", data)
	return data
}

func CreateFile(number int) {
	//membuat file
	file, err := os.Create("output.txt")

	if err != nil {
		log.Fatalf("failed creating file: %s", err)
	}

	//menutup file sebelum fungsi selesai dijalankan
	defer file.Close()
	strNumber := strconv.Itoa(number)
	len, err := file.WriteString("The sum is : " + strNumber)

	if err != nil {
		log.Fatalf("failed writing to file: %s", err)
	}

	fmt.Printf("File Name: %s\n", file.Name())
	fmt.Printf("Length: %d bytes\n", len)
}

func main() {
	input := ReadFile()
	result := sum(input)
	CreateFile(result)
}

//hasilnya bisa kita lihat di output.txt

func sum(input []byte) int {
	result := 0
	numbers, err := BytetoInt(input)
	if err != nil {
		fmt.Println("error :", err)
	}

	for _, number := range numbers {
		result += number
	}
	return result
}

func BytetoInt(bytes []byte) ([]int, error) {
	byteString := strings.Split(string(bytes), " ")
	result := []int{}
	for _, val := range byteString {
		integerFull, _ := strconv.Atoi(val)
		result = append(result, integerFull)
	}
	return result, nil
}
