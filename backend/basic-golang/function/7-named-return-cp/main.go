package main

import "fmt"

//fungsi square sama seperti sebelumnya
//yang membedakan adalah fungsi ini
//menggunakan named return
func main() {
	result1, result2 := square(4, 5)
	fmt.Printf("%d dan %d\n", result1, result2)
	fmt.Println(square(9, 8))
}

// TODO: answer here
func square (n1,n2 int)(result1,result2 int){
	 result1 = n1*n1
	 result2 = n2*n2
	return result1, result2
}




