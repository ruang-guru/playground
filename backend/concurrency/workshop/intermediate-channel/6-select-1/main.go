package main

import (
	"fmt"
	"time"
)

type person struct {
	name  string
	hobby string
}

//pada program ini terdapat person bernama john yang hobby mengerjakan soal matematika
//saat john bangun ia bisa makan dan minum
//untuk memberi makan dan minum, kita bisa mengirimkannya lewat channel foodInput dan drinkInput
//dia bisa memilih mau makan atau minum dulu
//tergantung dari makanan atau minumannya dulu yang siap
func main() {
	john := person{name: "john", hobby: "mengerjakan soal matematika"}
	foodInput := make(chan string)
	drinkInput := make(chan string)

	go john.awake(foodInput, drinkInput)
	foodInput <- "bakso"
	drinkInput <- "air"
	foodInput <- "sate"
	foodInput <- "nasi goreng"
	drinkInput <- "cola"
	foodInput <- "nasi padang"
	drinkInput <- "sirup"
	time.Sleep(100 * time.Millisecond)

}

func (p *person) awake(foodInput, drinkInput chan string) {
	for {
		select {
		case drink := <-drinkInput:
			fmt.Printf("%s minum %s\n", p.name, drink)
		case food := <-foodInput:
			fmt.Printf("%s makan %s\n", p.name, food)
		}
	}
}
