package main

import "fmt"

//Memanggil fungsi goodAfternoon
//dari dalam good afternoon akan dilakukan print "selamat sore name1 dan name2"
func main() {
	goodAfternoon("adi", "anti")
	goodAfternoon("ado", "suci")

}

// TODO: answer here
func goodAfternoon (name1 string, name2 string){
	fmt.Println("selamat sore ",name1,"dan ",name2)
}

