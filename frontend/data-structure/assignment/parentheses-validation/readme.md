# Parentheses Validation

Salah satu problem populer yang dapat diselesaikan dengan menggunakan Stack adalah mengecek validitas tanda kurung.
Diberikan sebuah string yang hanya terdapat karakter '(', ')', '{', '}', '[', dan ']'.
Tentukan apakah sebuah string merupakan sekumpulan tanda kurung yang valid.
String tanda kurung yang valid adalah

1. Tanda kurung buka harus ditutup oleh pasangannya.
2. Tanda kurung buka harus ditutup di urutan yang benar.

| Contoh ke | Input | Output | Penjelasan |
| --- | --- | --- | --- |
| 1 | "()" | true | tanda kurung buka '(' ditutup dengan pasangannya yaitu ')' |
| 2 | "{)" | false | tanda kurung buka '{' tidak ditutup oleh pasangannya |
| 3 | "([])" | true | tanda kurung buka ditutup dengan pasangannya dan sesuai dengan urutan |
