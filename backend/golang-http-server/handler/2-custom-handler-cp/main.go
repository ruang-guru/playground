package main

import (
	"fmt"
	"math/rand"
	"net/http"
)

// Dari contoh yang telah diberikan, buatlah sebuah custom handler yang dapat menampilkan random quotes.
// Buat handler QuotesHandler yang menampilkan quote dibawah ini secara acak.

var quotes = []string{
	"Be yourself; everyone else is already taken. ― Oscar Wilde",
	"Be the change that you wish to see in the world. ― Mahatma Gandhi",
	"I have not failed. I've just found 10,000 ways that won't work. ― Thomas A. Edison",
	"It is never too late to be what you might have been. ― George Eliot",
	"Everything you can imagine is real. ― Pablo Picasso",
	"Nothing is impossible, the word itself says 'I'm possible'! ― Audrey Hepburn",
}

// TODO: answer here

func (qh QuotesHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// TODO: answer here
}
