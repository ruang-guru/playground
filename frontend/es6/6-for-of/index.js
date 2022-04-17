/**
 * Menghitung berapa banyak huruf vokal di dalam sebuah kata
 *
 * Input: aku
 * Output: 2
 *
 * Catatan: gunakan for-of untuk perulangan, huruf vokal yang dihitung tidak boleh berubah (tidak distinct)
 */

const countVowels = (word) => {
  let totalVowels = 0;
  const vowels = ["a", "i", "u", "e", "o"];

  for (const letter of word) {
    // Perulangan for-of pada string
    for (const vowel of vowels) {
      // Perulangan for-of pada array
      if (letter === vowel) {
        totalVowels++;
      }
    }
  }
  return totalVowels;
};