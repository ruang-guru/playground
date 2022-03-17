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

	//beginanswer
	for {
		fmt.Println("Menu makanan:")
		i := 1
		for menu, price := range foodMenu {
			fmt.Println(i, ". Menu:", menu, ", ", "Price :", price)
			i++
		}
		fmt.Println()

		var menu, addAgain string
		for {
			fmt.Println("Form:")
			fmt.Printf("Masukan nama menu yang mau dipesan: ")
			fmt.Scan(&menu)

			if price, found := foodMenu[menu]; found {
				orderMenu[menu] = price
				break
			} else {
				fmt.Printf("%v tidak ditemukan!, Coba pilih menu lainnya.", menu)
			}
		}

		fmt.Println()
		fmt.Println("Menu telah ditambahkan ke Ordered Menu:")
		for menu, price := range orderMenu {
			fmt.Println("Menu:", menu, ", ", "Price :", price)
		}
		fmt.Println()

		fmt.Printf("Ingin memesan menu lain?(yes/no): ")
		fmt.Scan(&addAgain)

		if addAgain == "no" {
			fmt.Println("Data Order Updated:")
			for menu, price := range orderMenu {
				fmt.Println("Menu:", menu, ", ", "Price :", price)
			}

			break
		}

	}

	var total int64 = 0
	for _, val := range orderMenu {
		total += val
	}

	fmt.Println("Total harga makanan yang harus anda bayar: ", total)
	//endanswer
}
