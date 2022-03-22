package main

import "fmt"

func main() {
	// tambahkan setiap data pada slice dibawah ini
	data := []int{1, 20, 30, 55, 159, 100, 200, 300, 400, 500, 600, 700, 800, 900, 1000}

	// buatlah perulangan for loop dari data di atas

	// hint: gunakan len(data)
	result := 0
<<<<<<< HEAD

	for i := 0; i <= len(data)-1; i++ {
		// TODO: answer here
		result += data[i]

=======
	for i := 0; i <= len(data); i++ {
		// beginanswer
		result += data[i]
		// endanswer
>>>>>>> 0dbcdd8ebce63009fcee596516afbb40b893ca25
	}

	fmt.Println(result)

	// arr := []int{1, 2, 3}
	// res := 0
	// for i := 0; i < 3; i++ {
	// 	res += arr[i]
	// }

	// fmt.Println(res)
	// fmt.Println(len(arr))
}
