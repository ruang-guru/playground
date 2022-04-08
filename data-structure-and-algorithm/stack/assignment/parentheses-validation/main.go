package main

import "github.com/ruang-guru/playground/data-structure-and-algorithm/stack/assignment/parentheses-validation/stack"

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
	//beginanswer
	if len(s)%2 != 0 || len(s) == 0 {
		return false
	}

	parenthesesMap := map[rune]rune{
		'(': ')',
		'[': ']',
		'{': '}',
	}
	stack := stack.Stack{
		Top:  -1,
		Data: []rune{},
	}

	for _, c := range s {
		if closer, ok := parenthesesMap[c]; ok {
			stack.Push(closer)
			continue
		}

		if stack.IsEmpty() {
			return false
		}
		poppedValue, _ := stack.Pop()
		if poppedValue != c {
			return false
		}
	}

	return stack.IsEmpty()
	//endanswer
}
