/**
 * Muncul dimana-mana.
 * Merupakan nilai yang muncul di setiap subarray (Muncul dimana-mana).
 *
 * Contoh:
 *  - [[1, 2, 3], [5,4,1], [4, 14, 1]], 1 => true
 *  - [[1, 2, 3], [5,4,1], [4, 14, 1]], 2 => false
 *
 * Buatlah fungsi yang argumen pertamanya adalah sebuah array dua dimensi, dan argumen keduanya adalah elemen yang dicari. Jika value dari argumen ke dua terdapat pada setiap subarray maka akan mengembalikan nilai true, dan jika hanya terdapat pada 1 sub array maka akan mengembalikan nilai false.
 * Hint:
 * - Pastikan menggunakan for-of untuk perulangan
 */

const checkOmnipresent = (array, valueToFind) => {
  // beginanswer
  let numberOfPresent = 0;

  for (subarray of array) {
    for (value of subarray) {
      if (valueToFind === value) {
        numberOfPresent++;
      }
    }
  }
  return array.length === numberOfPresent;
  // endanswer
};

module.exports = checkOmnipresent