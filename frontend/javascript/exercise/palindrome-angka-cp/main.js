/**
 * Carilah angka palindrome yang LEBIH BESAR dan TERDEKAT dari angka-angka dibawah ini
 * 
 * Misal:
 * - Jika angka pada parameter bernilai 90 maka angka palindrome terdekat adalah 99
 * - Jika angka pada parameter bernilai 102 maka angka palindrome terdekat adalah 111
 */

function angkaPalindrome(num) {
  //beginanswer
  var cek = true;
  while (cek) {
    num++;
    var numStr = num.toString();
    var hasil = '';
    for (var j = numStr.length - 1; j >= 0; j--) {
      hasil += numStr[j];
    }
    if (hasil === numStr) {
      return num;
    }
  }
  //endanswer
}

console.log(angkaPalindrome(10)); //11
console.log(angkaPalindrome(17)); //22
console.log(angkaPalindrome(175)); //181

module.exports = angkaPalindrome