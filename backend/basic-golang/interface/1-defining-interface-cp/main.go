package main

import "fmt"

// Dari contoh yang telah diberikan, kamu dapat mencoba untuk membuat interface dan mengimplementasikan interface.
// Buatlah interface Employee yang memiliki method signature GetBonus() int
// Buatlah implementasi interface Employee pada objek Manager, SeniorEngineer, dan JuniorEngineer.
// Pada objek Manager, SeniorEngineer, dan JuniorEngineer memiliki satu atribut yaitu BaseSalary.
// Untuk menghitung bonus terdapat tiga aturan sebagai berikut:
// Bonus untuk Manager adalah 3 * BaseSalary
// Bonus untuk SeniorEngineer adalah 2 * BaseSalary
// Bonus untuk JuniorEngineer adalah 1 * BaseSalary

// TODO: answer here
type Employee interface {
	GetBonus() int
}

type Manager struct {
	BaseSalary int
}

func (m Manager) GetBonus() int {
	return 3 * m.BaseSalary
}

type SeniorEngineer struct {
	BaseSalary int
}

func (s SeniorEngineer) GetBonus() int {
	return 2 * s.BaseSalary
}

type JuniorEngineer struct {
	BaseSalary int
}

func (j JuniorEngineer) GetBonus() int {
	return j.BaseSalary
}

func TotalEmployeeBonus(employees []Employee) int {
	// Hitunglah total bonus yang dikeluarkan dari list of Employee
	// TODO: answer here
	total := 0
	for _, employee := range employees {
		total += employee.GetBonus()
	}

	return total
}

func main() {
	// Buatlah objek konkret untuk masing-masing objek dan panggil function TotalEmployeeBonus. Print total bonus untuk semua employee.
	// TODO: answer here

	manager, senior, junior := Manager{BaseSalary: 300000}, SeniorEngineer{BaseSalary: 150000}, JuniorEngineer{BaseSalary: 100000}
	employees := []Employee{manager, senior, junior}
	fmt.Println(TotalEmployeeBonus(employees))
}
