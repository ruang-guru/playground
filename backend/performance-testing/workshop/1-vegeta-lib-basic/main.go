package main

import (
	"fmt"
	"time"

	vegeta "github.com/tsenart/vegeta/v12/lib"
)

//jalankan main.go pada folder backend/performance-testing/server
//jalankan program ini

func main() {
	testGet()
	testPost()
}

func testGet() {
	duration := 1 * time.Second                            //durasi attack
	frequency := 40                                        //jumlah request
	target := "http://localhost:8090"                      //target
	rate := vegeta.Rate{Freq: frequency, Per: time.Second} //mengatur rate request
	targeter := vegeta.NewStaticTargeter(vegeta.Target{
		Method: "GET",
		URL:    target,
	}) //mengatur targeter vegeta
	metrics := vegetaAttack(targeter, rate, duration) //menjalankan vegeta attack
	fmt.Println(metrics.StatusCodes)                  //menampilkan status code
	fmt.Println(metrics.Latencies.Max)                // menampilkan latency maksimum
}

func testPost() {
	duration := 1 * time.Second
	frequency := 40
	target := "http://localhost:8090"
	body := "hello"
	rate := vegeta.Rate{Freq: frequency, Per: time.Second}
	targeter := vegeta.NewStaticTargeter(vegeta.Target{
		Method: "POST",
		URL:    target,
		Body:   []byte(body), //convert string ke byte agar bisa dikirim
	})
	metrics := vegetaAttack(targeter, rate, duration)
	fmt.Println(metrics.StatusCodes)
	fmt.Println(metrics.Latencies.Max)
}

func vegetaAttack(targeter vegeta.Targeter, rate vegeta.ConstantPacer, duration time.Duration) *vegeta.Metrics {
	attacker := vegeta.NewAttacker() //membuat attacker baru
	var metrics vegeta.Metrics       // menyiapkan metrics
	for res := range attacker.Attack(targeter, rate, duration, "Example") {
		metrics.Add(res) //menambahkan hasil attack ke dalam metrics
	} //melakukan vegeta attack
	metrics.Close()
	return &metrics
}
