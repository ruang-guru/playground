package main

import "fmt"

// Dari contoh yang telah diberikan, kamu dapat mencoba untuk membuat interface dan mengimplementasikan interface.
// Buatlah interface Employee yang memiliki method signature GetBonus() int
// Buatlah implementasi interface Employee pada objek Manager, SeniorEngineer, dan JuniorEngineer.
// Pada objek Manager, SeniorEngineer, dan JuniorEngineer memiliki satu atribut yaitu BaseSalary.
// Untuk menghitung bonus terdapat tiga aturan sebagai berikut:
// Bonus untuk Manager adalah 3 * BaseSalary
// Bonus untuk SeniorEngineer adalah 2 * BaseSalary
// Bonus untuk JuniorEngineer adalah 3 * BaseSalary

//beginanswer
type Employee interface {
	GetBonus() int
}

type Manager struct {
	BaseSalary int
}

func (m Manager) GetBonus() int {
	return m.BaseSalary * 3
}

type SeniorEngineer struct {
	BaseSalary int
}

func (m SeniorEngineer) GetBonus() int {
	return m.BaseSalary * 2
}

type JuniorEngineer struct {
	BaseSalary int
}

func (m JuniorEngineer) GetBonus() int {
	return m.BaseSalary
}

//endanswer

func TotalEmployeeBonus(employees []Employee) int {
	// Hitunglah total bonus yang dikeluarkan dari list of Employee
	//beginanswer
	total := 0
	for _, employee := range employees {
		total += employee.GetBonus()
	}

	return total
	//endanswer
}

func main() {
	// Buatlah objek konkret untuk masing-masing objek dan panggil function TotalEmployeeBonus. Print total bonus untuk semua employee.
	//beginanswer
	manager := Manager{
		BaseSalary: 20000000,
	}
	seniorEngineer := SeniorEngineer{
		BaseSalary: 15000000,
	}
	juniorEngineer := JuniorEngineer{
		BaseSalary: 10000000,
	}

	totalBonus := TotalEmployeeBonus([]Employee{manager, seniorEngineer, juniorEngineer})
	fmt.Println(totalBonus)
	//endanswer
}
