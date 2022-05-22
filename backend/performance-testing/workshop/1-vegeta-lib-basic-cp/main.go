package main

import (
	"time"

	vegeta "github.com/tsenart/vegeta/v12/lib"
)

//menjalankan vegeta attack dengan method GET
//durasi yang digunakan 2 detik dan rate/frequency 10
//target didapatkan dari parameter
func vegetaGet(target string) *vegeta.Metrics {
	// TODO: answer here
	duration := 2 * time.Second
	freq := 10
	rate := vegeta.Rate{
		Freq: freq, Per: time.Second,
	}
	targeter := vegeta.NewStaticTargeter(vegeta.Target{
		Method: "GET",
		URL:    target,
	})
	metrics := vegetaAttack(targeter, rate, duration)
	return metrics
}

//menjalankan vegeta attack dengan method POST
//durasi yang digunakan 2 detik dan rate/frequency 15
//target didapatkan dari parameter fungsi ini
func vegetaPost(target string) *vegeta.Metrics {
	// TODO: answer here
	duration := 2 * time.Second
	freq := 15
	rate := vegeta.Rate{
		Freq: freq, Per: time.Second,
	}
	targeter := vegeta.NewStaticTargeter(vegeta.Target{
		Method: "POST",
		URL:    target,
	})
	metrics := vegetaAttack(targeter, rate, duration)
	return metrics
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
