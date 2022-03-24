package main

import "fmt"

// Kita dapat membuat custom error dengan beberapa atribut yang kita perlukan.
// Dalam baris kode berikut, terdapat custom error ErrorDataNotFound dengan atribut message dan errCode.
type ErrorDataNotFound struct {
	message string
	errCode int32
}

// Karena error merupakan interface, kita harus mengimplementasikan method signature dari interface error yaitu Error() string.
// Kode dibawah ini merupakan implementasi dari method Error().
// Kita juga dapat menyesuaikan pesan error yang ingin di return.
func (e *ErrorDataNotFound) Error() string {
	return fmt.Sprintf("error %d: %s", e.errCode, e.message)
}

func GetAge(data map[string]int, name string) (int, error) {
	if _, ok := data[name]; !ok {

		// Pada baris berikut terdapat return error dari custom error yang telah dibuat.
		return 0, &ErrorDataNotFound{
			message: fmt.Sprintf("%s not found in data", name),
			errCode: 404,
		}
	}

	return data[name], nil
}

func main() {
	peopleAge := map[string]int{
		"John": 20,
		"Raz":  8,
		"Tony": 40,
	}

	_, err := GetAge(peopleAge, "Roger")
	if err != nil {
		fmt.Println(err.Error())
	}
}
