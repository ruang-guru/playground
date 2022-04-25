package main

import (
	"fmt"
	"time"
)

type person struct {
	name  string
	hobby string
}

//hmm bagaimana jika john sedang tidak lagi makan dan minum
//kemungkinan dia akan melakukan hobbynya
//dengan select, kita bisa menambahkan default
//default akan dijalankan ketika case lain belum ada yang siap (belum ada data yang dikirim ke channel)
func main() {
	john := person{name: "john", hobby: "mengerjakan soal matematika"}
	foodInput := make(chan string)
	drinkInput := make(chan string)

	go john.awake(foodInput, drinkInput)
	foodInput <- "bakso"
	foodInput <- "sate"
	foodInput <- "nasi goreng"
	foodInput <- "nasi padang"
	drinkInput <- "air"
	drinkInput <- "cola"
	drinkInput <- "sirup"

}

func (p *person) awake(foodInput, drinkInput chan string) {
	for {
		select {
		case drink := <-drinkInput:
			fmt.Printf("%s minum %s\n", p.name, drink)
		case food := <-foodInput:
			fmt.Printf("%s makan %s\n", p.name, food)
		default:
			fmt.Printf("%s %s\n", p.name, p.hobby)
			time.Sleep(100 * time.Millisecond) // mengerjakan soal matematika
		}
	}
}
