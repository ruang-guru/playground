package main

import "fmt"

// Check Point:
// Cashier
// - Input: Nama Menu, Add Again
// - Output: Total Harga yang harus dibayar

// Contoh:
// Menu makanan:
// 1 . Menu: bakso ,  Price : 20000
// 2 . Menu: burger ,  Price : 15000
// 3 . Menu: sate ,  Price : 25000
// 4 . Menu: sosis ,  Price : 20000
// 5 . Menu: soto ,  Price : 25000

// Input:
// - Form:
//   - Masukan nama menu yang mau dipesan: bakso

// Output:
// Menu telah ditambahkan ke Ordered Menu:
// Menu: bakso ,  Price : 20000

// Input:
// - Ingin memesan menu lain?(yes/no): if no (break) --> if yes (add again)

// Output:
// Data Order Updated:
//  - Menu: sate ,  Price : 25000
//  - Menu: burger ,  Price : 15000
//  - Menu: sosis ,  Price : 20000
//  - Total harga makanan yang harus anda bayar:  60000

func main() {
	foodMenu := map[string]int64{
		"bakso":  20000,
		"burger": 15000,
		"sate":   25000,
		"sosis":  20000,
		"soto":   25000,
	}

	orderMenu := make(map[string]int64)

	// TODO: answer here
	for {
		fmt.Println("======Menu=====")
		for food, price := range foodMenu {
			fmt.Println("1. Menu: ", food, ", price: ", price)
		}
		fmt.Println(" ")
		var food, addAgain string
		var price int

		fmt.Println("Form Pemesanan: ")
		fmt.Printf("Menu yang dipesan: ")
		fmt.Scan(&food)
		fmt.Printf("Pesan lagi?(yes/no)")
		fmt.Scan(&addAgain)

		if addAgain == "no" {
			break
		}
	}
	fmt.Println("Menu yang kamu pesan: ")
	for food, price := orderMenu{
		fmt.Println("-", food, "harga", price)
		fmt.Println("Total: Rp")
	}

}
