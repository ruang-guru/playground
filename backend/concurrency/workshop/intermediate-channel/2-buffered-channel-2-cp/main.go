package main

import (
	"fmt"
	"time"
)

var person = "andi"
var names = []string{"budi", "toni", "adi", "ado", "alif", "yudi"}

//mengembalikan string, dimana `name` menyapa semua `names`
func greetAll(person string, names []string, output chan<- string) {
	// TODO: answer here
	fmt.Println("selesai mengirim")

}

// buat size buffered channel sesuai jumlah names
func testBufferedChannel(result chan<- string) {
	output := make(chan string) // TODO: replace this

	go greetAll(person, names, output)
	time.Sleep(100 * time.Millisecond)
	for i := 0; i < 6; i++ {
		greeting := <-output
		result <- greeting
	}
}

//goroutine greetAll dapat mengirim ke goroutine testBufferedChannel walaupun channel belum siap menerima
