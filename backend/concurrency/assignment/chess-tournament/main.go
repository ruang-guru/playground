package main

import (
	"fmt"
	"math/rand"
	"time"
)

var start time.Time

func init() {
	start = time.Now()
}

// Pak Wayne sedang mengadakan turnamen catur tingkat RT dengan total 10 match
// Saat ini rencana Pak Wayne adalah menjalankan satu match dan menunggu match itu selesai
// Pak Wayne juga ingin tau jika ada match yang mulai atau selesai
// Kamu dimintai tolong mencarikan solusi agar turnamen dapat selesai lebih cepat
// Apa yang kamu lakukan ? jawab di assignment-2

func main() {
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < 10; i++ {

		fmt.Printf("match #%d started\n", i+1)
		time.Sleep(100 * time.Millisecond)
		fmt.Printf("match #%d finished\n", i+1)

	}
	fmt.Println("all the match finished. Total time needed:", time.Since(start))
}
