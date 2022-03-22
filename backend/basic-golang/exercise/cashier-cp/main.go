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

<<<<<<< HEAD
	// TODO: answer here
	fmt.Println("\n\nMenu Makanan")
	i := 1
	var pesan, reOrder string
	for menu, harga := range foodMenu {
		fmt.Printf("%d . Menu : %s , price %d\n", i, menu, harga)
		i++
	}

	for {
		fmt.Print("\nMasukan nama menu yang akan dipesan : ")
		fmt.Scan(&pesan)

		// cek apakah user memasukan menu dengan benar
		_, isAvailable := foodMenu[pesan]
		// jika benar maka
		if isAvailable {
			orderMenu[pesan] = foodMenu[pesan]
			fmt.Println("\nMenu telah ditambahkan ke Ordered Menu")
			iter := 1

			for menu, price := range orderMenu {
				fmt.Printf("%d . Menu : %s, Price : %d \n", iter, menu, price)
				iter++
			}

			fmt.Print("\nIngin memesan menu lain ? (yes/no)")
			fmt.Scan(&reOrder)

			if reOrder == "NO" || reOrder == "no" || reOrder == "No" || reOrder == "nO" {
				fmt.Println("\n\nData Order Updated")
				iterasi := 1
				total := 0
				for menu, harga := range orderMenu {
					fmt.Printf("%d . Menu : %s, Price : %d\n", iterasi, menu, harga)
					iterasi++
					total += int(harga)
				}
				fmt.Printf("Total harga yang harus anda bayar : %d\n", total)
				break
			}

		} else {
			fmt.Printf("\nMaaf menu %s tidak ada dalam daftar!!!\n", pesan)
=======
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
>>>>>>> 0dbcdd8ebce63009fcee596516afbb40b893ca25
		}

	}

<<<<<<< HEAD
=======
	var total int64 = 0
	for _, val := range orderMenu {
		total += val
	}

	fmt.Println("Total harga makanan yang harus anda bayar: ", total)
	//endanswer
>>>>>>> 0dbcdd8ebce63009fcee596516afbb40b893ca25
}
