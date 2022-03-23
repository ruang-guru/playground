// Tulis kode untuk menjumlahkan semua nilai dan simpan dalam variabel sum
// Jika nilai kosong, maka hasilnya harus 0.
//
// contoh: 
// let scores = {
//    Andi: 100,
//    Budi: 70,
//    Lisa: 75
//  }
//
// result: 245

let scores = {
    Andi: 100,
    Budi: 70,
    Lisa: 75
};
  
let sum = 0;
for (let key in scores) {
  sum += scores[key];
}
  
console.log(sum);