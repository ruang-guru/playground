package main

import (
	"fmt"
	"time"
)

var start time.Time

func init() {
	start = time.Now()
}

//kode di atas hanya untuk membantu mengetahui waktu program dijalankan

func greetAndy(called *bool) {
	time.Sleep(100 * time.Millisecond)
	*called = true
	fmt.Println("hi Andy. at time", time.Since(start))
}

func greetBob(called *bool) {
	time.Sleep(100 * time.Millisecond)
	*called = true
	fmt.Println("hi Bob. at time", time.Since(start))
}

//Bagaimana cara agar greetAndy dan greetBob dapat berjalan secara concurrent ?
func call(calledAndy, calledBob *bool) {
	// TODO: answer here
	greetAndy(calledAndy)
	// TODO: answer here
	greetBob(calledBob)
	time.Sleep(200 * time.Millisecond)
	fmt.Println("from call function at time", time.Since(start))
	fmt.Println("called", *calledAndy && *calledBob)
}
