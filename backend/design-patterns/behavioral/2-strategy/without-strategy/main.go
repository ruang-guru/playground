package main

import "fmt"

// implementasi sederhana untuk membahas strategy pattern
// kita membuat toko online yang melayani tiga metode pembayaran

type Payment string

const (
	Bank   Payment = "bank"
	EMoney Payment = "e-money"
	COD    Payment = "cod"
)

type ShoppingCart struct {
	Payment Payment
}

func NewShoppingCart() *ShoppingCart {
	return &ShoppingCart{}
}

func (i *ShoppingCart) ChoosePayment(p Payment) {
	i.Payment = p
}

func (i *ShoppingCart) Pay() {
	//asumsi semua id dan pass benar serta uang cukup
	if i.Payment == Bank {
		fmt.Println("using Bank")
		var id string
		fmt.Print("enter your id: ")
		fmt.Scanln(&id)
		var password string
		fmt.Print("enter your password: ")
		fmt.Scanln(&password)
	} else if i.Payment == EMoney {
		//asumsi username sudah ready
		fmt.Println("using E-Money")
		var password string
		fmt.Print("enter your password: ")
		fmt.Scanln(&password)
	} else if i.Payment == COD {
		fmt.Println("using COD")
		//asumsi sudah ada alamat COD jadi tidak perlu melakukan apa"
		fmt.Println("using address that available in your profile")
	}
	fmt.Println("SUCCESS")
	fmt.Println("--------------------------")
}

func main() {
	cart := NewShoppingCart()
	fmt.Println("choose payment method")
	fmt.Println("1.Bank\n2.E-Money\n3.COD")
	var input int
	fmt.Scanln(&input)
	if input == 1 {
		cart.Payment = Bank
	} else if input == 2 {
		cart.Payment = EMoney
	} else if input == 3 {
		cart.Payment = COD
	}
	cart.Pay()
}

//kita bisa menggunakan cara ini untuk menyelesaikan proses pembayaran
//tapi jika kita melihat method `Pay()`, baris kode pada method tersebut cukup panjang
//dan jika kita perlu menambahkan metode pembayaran baru seperti `PayLater`
//kita jadi perlu mengubah method `Pay()` dan baris kodenya akan menjadi semakin panjang
//kenapa perlu menghindari mengubah method `Pay()` ?
//jika kita melakukan kesalahan penulisan, dampaknya bisa kena ke payment lain
//kita bisa menggunakan `strategy pattern` untuk menyelsaikan hal ini
