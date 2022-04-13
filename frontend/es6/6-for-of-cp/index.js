/**
 * Tampilkan tipe data dari setiap elemen array
 * Input: ["A", 3, [2]]
 * Outout: ["string", "number", "object"]
 *
 * Hint:
 * - Untuk pengecekan tipe data kalian dapat menggunakan typeof
 * - Pastikan menggunakan for-of untuk perulangan
 */

const convertElementToType = (array) => {
  // beginanswer
  const newArray = [];
  for (element of array) {
    newArray.push(typeof element);
  }
  return newArray;
  // endanswer
};

module.exports = convertElementToType