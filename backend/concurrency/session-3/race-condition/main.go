package main

import (
	"fmt"
	"time"
)

func main() {
	a := 0
	for i := 0; i < 1000; i++ {
		go func() {
			a++
		}()
	}
	time.Sleep(500 * time.Millisecond)
	fmt.Println(a) // apa nilai dari variabel a?
}

// Berapa kah nilai akhir a? Kita tidak tahu (undeterministic)
//tergantung bagaimana timing threads tersebut interleaved.
