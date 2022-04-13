/*
* Levi baru membuka Cafe bernama Ruang Cafe. saat ini cafe yang dimiliki Levi hanya memiki menu juice dan latte.
* Harga untuk menu juice yaitu Rp. 50.000 sedangkan untuk latte Rp. 300.000.
* Levi ingin membuat sebuah program untuk melihat menu apa yang cocok untuk pelanggannya.
* Ada pun kriteria dari aplikasi yang ingin dibuat yaitu:
* - jika pelanggan tidak menginput nama maka akan nemapilkan "Anda tidak boleh masuk!"
* - jika uang yang dimiliki oleh pelanggan kurang maka akan menampilkan "Uang tidak cukup. Anda harus pulang."
* - jika uang pelanggan cukup dan usia dibawah 17 tahun maka akan menampilkan "Anda bisa pesan juice. Sisa uang anda: <jumlah uang tersisa>"
* - jika uang pelanggan cukup dan usia 17 tahun keatas maka akan menampilkan "Anda bisa pesan latte. Sisa uang anda: <jumlah uang tersisa>"
* Bantulah Levi untuk membuat Pseudocode beserta Codenya.
* 
* notes:
* - Dilarang MENGGANTI nama variable yang sudah diberikan
* - Terdapat variable juice dan latte
*/


// PSEUDOCODE :
//beginanswer
// STORE "juice" with 50000
// STORE "latte" with 300000
//
// IF "name" LENGTH === 0
//   RETURN "Anda tidak boleh masuk!"
// IF "age" < 17
//   IF "money" > "juice"
//    CALCULATE "money" minus "juice"
//    SET sisa uang value with calculation result
//    RETURN "Anda bisa pesan juice. Sisa uang anda: "
//   ELSE 
//    RETURN "Uang tidak cukup. Anda harus pulang."
// IF "age" > 17
//   IF "money" > "latte"
//    CALCULATE "uang" minus "latte"
//    SET sisa uang value with calculation result
//    RETURN "Anda bisa pesan latte. Sisa uang anda: "
//   ELSE 
//    RETURN "Uang tidak cukup. Anda harus pulang."
//endanswer

function ruangCafe(name, age, money) {
  const juice = 50000;
  const latte = 300000;

  //beginanswer
  if (name.length === 0) {
    return "Anda tidak boleh masuk!"
  }
  else {
    if (age < 17) {
      if (money >= juice) {
        return `Anda bisa pesan juice. Sisa uang anda: ${money - juice}`
      }
      else {
        return "Uang tidak cukup. Anda harus pulang."
      }
    }
    else {
      if (money >= latte) {
        return `Anda bisa pesan latte. Sisa uang anda: ${money - latte}`
      }
      else {
        return "Uang tidak cukup. Anda harus pulang."
      }
    }
  }
  //endanswer
}

console.log(ruangCafe('', 21, 2000000))
console.log(ruangCafe('Dito', 16, 1000000))
console.log(ruangCafe('Adit', 11, 9000))
console.log(ruangCafe('Fauzan', 30, 999999))
console.log(ruangCafe('Tegar', 27, 5000))

module.exports = ruangCafe