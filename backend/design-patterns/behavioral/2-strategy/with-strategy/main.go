package main

import "fmt"

//kita ubah type string sebelumnya menjadi interface (strategy)
type Payment interface {
	Pay()
}

type ShoppingCart struct {
	Payment Payment
}

func NewShoppingCart() *ShoppingCart {
	return &ShoppingCart{}
}

func (i *ShoppingCart) ChoosePayment(p Payment) {
	i.Payment = p
}

//selanjutnya kita ubah Pay() agar menggunakan method dari Payment (strategy) yang dipilih
func (i *ShoppingCart) Pay() {
	//asumsi semua id dan pass benar serta uang cukup
	i.Payment.Pay()
	fmt.Println("SUCCESS")
	fmt.Println("--------------------------")
}

func main() {
	cart := NewShoppingCart()
	fmt.Println("choose payment method")
	fmt.Println("1.Bank\n2.E-Money\n3.COD\n4.PayLater")
	var input int
	fmt.Scanln(&input)
	if input == 1 {
		cart.Payment = &Bank{}
	} else if input == 2 {
		cart.Payment = &EMoney{}
	} else if input == 3 {
		cart.Payment = &COD{}
	} else if input == 4 {
		cart.Payment = &PayLater{}
	}
	cart.Pay()
}

//kebawah adalah implementasi konkrit dari tiap strategy

type Bank struct{}

func (b *Bank) Pay() {
	fmt.Println("using Bank")
	var id string
	fmt.Print("enter your id: ")
	fmt.Scanln(&id)
	var password string
	fmt.Print("enter your password: ")
	fmt.Scanln(&password)
}

type EMoney struct{}

func (EM *EMoney) Pay() {
	//asumsi username sudah ready
	fmt.Println("using E-Money")
	var password string
	fmt.Print("enter your password: ")
	fmt.Scanln(&password)
}

type COD struct{}

func (C *COD) Pay() {
	fmt.Println("using COD")
	//asumsi sudah ada alamat COD jadi tidak perlu melakukan apa"
	fmt.Println("using address that available in your profile")
}

//jika ingin menambahkan payment (strategy) baru kita cukup membuat
//implementasi interface baru dan mengubah client code (main)
//dengan ini kita tidak perlu selalu mengubah method `Pay`

type PayLater struct{}

func (p *PayLater) Pay() {
	var password string
	fmt.Print("enter your PIN: ")
	fmt.Scanln(&password)
	fmt.Println("sending the bill to your account")
}
