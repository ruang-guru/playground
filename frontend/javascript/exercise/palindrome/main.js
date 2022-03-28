// cek apakah kata yang dimasukan merupakan kata palindrome.

function palindrome(kata) {
  var hasil = '';
  for (var i = kata.length - 1; i >= 0; i--) {
    hasil += kata[i];
  }
  if (kata === hasil) {
    return true
  } else {
    return false
  }
}

console.log(palindrome('katak')); //true
console.log(palindrome('kasur rusak')); //true
console.log(palindrome('makan'));  //false