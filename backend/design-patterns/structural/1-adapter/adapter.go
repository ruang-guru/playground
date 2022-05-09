package main

// Coba katakana pada saat kalian ingin membuat aplikasi
// Kalian tidak bisa mengutak atik kode dari turkey.go

// Nah yang kita lakukan adalah mendefine sebuah adapter

type TurkeyAdapter struct {
	Turkey Turkey
}

func (t *TurkeyAdapter) Quack() {
	t.Turkey.Gobble()
}

func (t *TurkeyAdapter) Fly() {
	for i := 0; i < 5; i++ {
		t.Turkey.Fly()
	}
}
