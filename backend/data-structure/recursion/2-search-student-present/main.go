package main

import (
	"fmt"
)

func searchStudentPresent(students []string, siswa string, len int) string {
	var absensi string
	if len == -1 {
		absensi = siswa + " tidak hadir"
	} else if students[len] == siswa {
		absensi = siswa + " hadir"
	} else {
		return searchStudentPresent(students, siswa, len-1)
	}
	return absensi
}

func main() {
	studentsPresent := []string{"Dino", "Gilang", "Rangga", "Baren", "Dedi", "Dewi", "Juan"}

	//Check if student is present
	var student string
	fmt.Print("Masukkan nama siswa yang ingin dicari : ")
	fmt.Scanln(&student)

	result := searchStudentPresent(studentsPresent, student, len(studentsPresent)-1)
	fmt.Println(result)
}
