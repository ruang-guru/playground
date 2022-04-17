//sudah ditambahkan fitur untuk mengecek waktu berjalan
package main

import (
	"fmt"
	"runtime"
	"time"
)

var start time.Time

func init() {
	start = time.Now()
}

func main() {
	runtime.GOMAXPROCS(1) //agar berjalan parallel hapus ini
	fmt.Println("Main started", time.Since(start))

	go func() {
		// time.Sleep(11 * time.Millisecond)
		fmt.Println("Hello from new goroutine", time.Since(start))
	}()
	fmt.Println("Before sleep", time.Since(start))
	time.Sleep(10 * time.Millisecond) // menunggu 10 milli-second agar go routine berjalan
	fmt.Println("Main stopped", time.Since(start))
}
