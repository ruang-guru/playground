//jangan ditujunkan ke peserta
package main

import "fmt"

func main() {
	hello := func(name string) func(string) string {
		return func(title string) string {
			return fmt.Sprintf("halo %s %s", title, name)
		}
	}

	helloDante := hello("dante")
	helloJack := hello("jack")

	fmt.Println(helloDante("mas"))
	fmt.Println(helloJack("bang"))

}
