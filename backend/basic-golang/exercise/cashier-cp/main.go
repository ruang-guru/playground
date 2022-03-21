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

	// TODO: answer here
	for {
		for name, price := range foodMenu {
			fmt.Println("Menu: "+name+",  Price:", price)
		}

		var order string
		fmt.Print("Masukan nama menu yang mau dipesan: ")
		fmt.Scan(&order)

		for name, price := range foodMenu {
			if name == order {
				orderMenu[name] = price
			}
		}

		fmt.Println()
		
		fmt.Println("Menu telah ditambahkan ke Ordered Menu:")
		for name, price := range orderMenu {
			fmt.Println("Menu: "+name+",  Price:", price)
		}

		fmt.Println()

		var again string
		fmt.Print("Ingin memesan menu lain?(yes/no) ")
		fmt.Scan(&again)

		fmt.Println()

		if (again == "no") {
			var total int64

			for name, price := range orderMenu {
				fmt.Println("Menu: "+name+",  Price:", price)

				total += price
			}
			fmt.Println("Total harga makanan yang harus anda bayar:", total)

			break
		}
	}
}
