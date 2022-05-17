package perusahaan

// Setiap perusahaan kan mempunyai 'Employee'
// Let say kita ingin membuat sistem dimana kita bisa mendapatkan total salary dari sebuah Group yang dibawahi oleh CTO atau VP

type Employee interface {
	GetSalary() int
	TotalDivisonSalary() int
}
