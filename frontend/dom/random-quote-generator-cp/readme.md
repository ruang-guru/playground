# Random Quote Generator

Klik gambar untuk video demo aplikasi.
[![](./assets/img.png)](https://drive.google.com/file/d/1yZpxloVzFp6sbPo5cOMVLBjbPd2NEsks/view?usp=sharing)

Random Quote Generator merupakan sebuah aplikasi untuk mengenerate random quote.
Setiap kali button "Generate Quote" di klik akan megenerate random quote.
Terdapat perbedaan pada setiap data quote tersedia.
Kalian diminta untuk menampilkan sebuah random quote, jika tidak tedapat citation dan year pada quote yang didapat maka kalian hanya diminta untuk menampilkan qoute beserta author nya.

Terdapat beberpa variable dan function dengan masing-masing fungsi yang berbeda.

## Variable

- `quote` = menangkap sebuah element untuk menampilkan quote
- `author` = menangkap sebuah element untuk menampilkan author
- `citation` = menangkap sebuah element untuk menampilkan citation
- `year` = menangkap sebuah element untuk menampilkan year
- `button` = menangkap sebuah element untuk melakuan generate random quote


## Function

- `getQuote` = function yang bertugas untuk mendapatkan 1 random quote dari daftar quote yang tersedia.
- `displayQuote` = function yang bertugas untuk menampilkan quote yang didapat pada function `getQuote`.

## Submission

`grader-cli submit cypress/integration/dom/random-quote-generator`