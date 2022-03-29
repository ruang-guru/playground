package main

import "github.com/ruang-guru/playground/backend/basic-golang/package/1-create-package/person"

// Package biasanya terletak di dalam folder tersendiri. Untuk membuat nama dari package kita dapat menggunakan beberapa aturan berikut:
// - Menggunakan terminologi yang singkat dan jelas.
// - Menggunakan lower case.
// - Umumnya menggunakan kata benda.
// - Tidak menggunakan snake case dan camel case.

// Sebagai contoh, dalam folder ini terdapat folder person. Di dalam folder person terdapat file person.go dan memiliki nama package person.
// Untuk menggunakan kode di dalam package person diluar package person, maka kita perlu melakukan import.
// Dalam kasus ini, pada package main kita akan menggunakan package person, membuat objek dari Person dan memanggil method SayHello().

func main() {
	tony := person.Person{
		Name: "Tony",
		Age:  40,
	}
	tony.SayHello()
}
