package main

import (
	"fmt"

	"github.com/ruang-guru/playground/backend/data-structure/assignment/parentheses-validation/stack"
)

// Salah satu problem populer yang dapat diselesaikan dengan menggunakan Stack adalah mengecek validitas tanda kurung.
// Diberikan sebuah string yang hanya terdapat karakter '(', ')', '{', '}', '[', dan ']'.
// Tentukan apakah sebuah string merupakan sekumpulan tanda kurung yang valid.
// String tanda kurung yang valid adalah
// 1. Tanda kurung buka harus ditutup oleh pasangannya.
// 2. Tanda kurung buka harus ditutup di urutan yang benar.

// Contoh 1
// Input: s = "()"
// Output: true
// Penjelasan: tanda kurung buka '(' ditutup dengan pasangannya yaitu ')'.

// Contoh 2
// Input: s = "{)"
// Output: false
// Penjelasan: tanda kurung buka '{' tidak ditutup oleh pasangannya.

// Contoh 3
// Input: s = "([])"
// Output: true
// Penjelasan: tanda kurung buka ditutup dengan pasangannya dan sesuai dengan urutan.

func IsValidParentheses(s string) bool {
	// TODO: answer here
	if len(s)%2 == 1 {
		return false
	}

	stackString := stack.Stack{
		Top:  -1,
		Data: []rune{},
	}

	isValid := false
	for _, val := range s {
		if string(val) == ")" || string(val) == "}" || string(val) == "]" {
			dataPop, err := stackString.Pop()
			if err != nil {
				fmt.Println(err.Error())
				return false
			}
			if (string(dataPop) == "(" && string(val) == ")") || (string(dataPop) == "{" && string(val) == "}") || (string(dataPop) == "[" && string(val) == "]") {
				isValid = true
				continue
			} else {
				isValid = false
				break
			}
		} else {
			stackString.Push(val)
		}

		isValid = false
	}

	return isValid
}
