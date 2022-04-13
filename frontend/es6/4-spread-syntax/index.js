/**
 * 1. Fungsi yang menerima 2 array sebagai argumen dan menyelipkan array kedua di tengah-tengah array pertama.
 *    Array pertama akan selalu memiliki dua element.
 */
function tuckIn(arr1, arr2) {
  // Memastika arr1 hanya memiliki 2 element.
  if (arr1.length > 2) {
    throw new Error("First argument has too many elements");
  }

  // menyisipkan array kedua di antara array pertama dengan menggunakan spread syntax.
  return [arr1[0], ...arr2, arr1[1]];
}

/**
 * Penggunaan spread syntax pada argumen sebuah fungsi
 */

function sum(a, b, c) {
  return a + b + c;
}

function sumArr(arr) {
  // Array kemudian di ekspand kedalam sebuah argumen.
  return sum(...arr);
}

/**
 * 3. Contoh senderhana menggabungkan 2 objek dengan menggunakan spread syntax
 */
function shallowMerge(obj1, obj2) {
  return { ...obj1, ...obj2 };
}
