"use strict";

/**
 * 1. Fungsi yang menerima input argumen berupa sembarang array,
 *    dan Mengembalikan array tipe data dari setiap elemen array input.
 **/
const getValueType = (arr) => arr.map((value) => typeof value);

/**
 * 2. Filter array angka, kembalikan angka ganjil saja
 * Jika menemukan bukan angka kembalikan -
 */
const filterOddNUmber = (numbers) =>
  numbers.filter((number) => {
    // Kembalikan "-" jika bukan angka
    if (isNaN(number)) return "-";

    // Jika tidak habis dibagi 2 maka ganjil.
    return number % 2 > 0;
  });

/**
 * 3. Fungsi yang menerima array kemudian menjumlahkan semua nilai dan mengembalikan angka
 *    Jika dalam elemen ditemukan bukan angka maka tukar dengan angka 0
 */
const sumArray = (numbers) =>
  numbers.reduce((previous, current) => {
    // Jika bukan angka tukar dengan angka 0
    const value = isNaN(current) ? 0 : current;

    // Tambahkan dengan hasil perulangan sebelumnya.
    return previous + value;
  }, 0);

