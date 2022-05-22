package main

import (
	"encoding/json"
	"fmt"

	// TODO: answer here
	"time"

	vegeta "github.com/tsenart/vegeta/v12/lib"
)

type Movie struct {
	ID      int    `json:"id"`
	Episode int    `json:"episode"`
	Name    string `json:"name"`
}

//Baca README untuk tau jumlah request yang perlu dilakukan dan targetnya
//untuk durasi cukup gunakan satu detik

//menambahkan movie baru
//untuk data yang dikirim adalah JSON
//gunakan struct Movie diatas, cukup gunakan field episode dan name
//ID sudah auto increment
func addMovieTest(target string) *vegeta.Metrics {
	// TODO: answer here
	newMovie := Movie{
		Episode: 1,
		Name:    "New Movie",
	}
	marshalJson, _ := json.Marshal(newMovie)

	targeter := vegeta.NewStaticTargeter(vegeta.Target{
		Method: "POST",
		URL:    target,
		Body:   marshalJson,
	})
	metrics := vegetaAttack(targeter, 10, time.Second)
	return metrics
}

//mendapatkan informasi movie dengan ID 1-25
//vegeta.NewStaticTargeter() adalah variadic function
//kita bisa menggunakannya untuk menentukan multiple target vegeta attack
func getMovieTest(target string) *vegeta.Metrics {
	// TODO: answer here
	targets := []vegeta.Target{}
	for i := 1; i <= 25; i++ {
		targets = append(targets, vegeta.Target{
			Method: "GET",
			URL:    fmt.Sprintf("%s/movie/%d", target, i),
		})
	}

	targeter := vegeta.NewStaticTargeter(targets...)
	metrics := vegetaAttack(targeter, 25, time.Second)
	return metrics
}

//mendapatkan semua informasi movie
func getMoviesTest(target string) *vegeta.Metrics {
	// TODO: answer here
	targeter := vegeta.NewStaticTargeter(vegeta.Target{
		Method: "GET",
		URL:    target,
	})
	metrics := vegetaAttack(targeter, 20, time.Second)
	return metrics
}

func vegetaAttack(targeter vegeta.Targeter, frequency int, duration time.Duration) *vegeta.Metrics {
	rate := vegeta.Rate{Freq: frequency, Per: time.Second}
	attacker := vegeta.NewAttacker()
	var metrics vegeta.Metrics
	for res := range attacker.Attack(targeter, rate, duration, "Example") {
		metrics.Add(res)
	}
	metrics.Close()
	return &metrics
}
