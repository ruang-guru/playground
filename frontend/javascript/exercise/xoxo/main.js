/*
Hitung dan cek lah jumlah karakter dari masing-masing huruf "x" dan "o".
Jika jumlah karakter "x" dan "o" sama maka akan mengembalikan "true" dan jika berbeda maka akan mengembalikan "false".
*/

function cekJumlahHuruf(kata) {
  var tipe1 = 0;
  var tipe2 = 0;
  for (var i = 0; i < kata.length; i++) {
    if (kata[i] === "x") {
      tipe1 += 1
    }
    else if (kata[i] === "o") {
      tipe2 += 1
    }
  }
  if (tipe1 === tipe2) {
    return true
  }
  else {
    return false
  }
}

console.log(cekJumlahHuruf("xoxxooox")); // true
console.log(cekJumlahHuruf("xoxoxoxo")); // true
console.log(cekJumlahHuruf("xxxoxox")); // false
console.log(cekJumlahHuruf("ooxoxxooo")); // false
console.log(cekJumlahHuruf("ooooxxoxxoxx")); // true