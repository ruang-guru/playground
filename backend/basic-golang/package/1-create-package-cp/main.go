package main

import (
	"fmt"

	"github.com/ruang-guru/playground/backend/basic-golang/package/1-create-package-cp/account"
)

// Dari contoh yang telah diberikan, buatlah sebuah package baru dengan nama account.
// Di dalam package tersebut, implementasikan method GetBalance yang berfungsi untuk mengecek saldo/balance
// dan implementasikan method Deposit yang berfungsi untuk menambah nilai pada saldo/balance

func main() {
	danaDarurat := account.Account{
		Name:    "Dana Darurat",
		Balance: 200,
	}
	fmt.Println(danaDarurat.GetBalance())

	danaDarurat.Deposit(100)
	fmt.Println(danaDarurat.GetBalance())
}
