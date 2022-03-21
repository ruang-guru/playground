package main

// Terdsapat tiga 3 jenis control structure yang ada pada Go

/*
	if-else two-way conditional execution block.
	for loop block.
	switch-case multi-way conditional execution block.
*/

func main() {
	var a int = 10
	var b int = 20

	var max int
	if a > b {
		max = a
	} else {
		max = b
	}
	print(max)
}
