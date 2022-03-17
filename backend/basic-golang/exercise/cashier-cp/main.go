package main

import (
	"fmt"
)

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

	i := 1

	for menu, price := range foodMenu {
		fmt.Printf("%d. Menu: %s , Price : %d\n", i, menu, price)
		i++
	}

	// TODO: answer here
	for {
		var order string
		fmt.Print("Masukan nama menu yang mau dipesan: ")
		fmt.Scan(&order)

		if val, ok := foodMenu[order]; ok {
			fmt.Println("Menu telah ditambahkan ke Ordered Menu:")
			fmt.Printf("Menu: %s , Price : %d \n", order, val)
			orderMenu[order] = val
		} else {
			fmt.Println("Menu tidak ditemukan")
		}

		var addAgain string
		fmt.Print("Ingin memesan menu lain?(yes/no): ")
		fmt.Scan(&addAgain)
		if addAgain == "no" {
			break
		}
	}

	fmt.Println("Data Order Updated:")
	var total int64 = 0
	for menu, price := range orderMenu {
		fmt.Printf("- Menu: %s , Price : %d \n", menu, price)
		total += price
	}
	fmt.Print("- Total harga makanan yang harus anda bayar:  ", total)
}
