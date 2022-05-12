// Mengembalikan pola * yang menurun di setiap barisnya
//
// contoh:
// baris = 5
// hasil:
// *****
// ****
// ***
// **
// *

// Masukan jumlah baris
const n = parseInt(prompt("Masukan jumlah baris: "));

// TODO: answer here
for (let i = 1; i <= n; i++) {
  let bintang = "";
  for (let j = n; j >= i; j--) {
    bintang = bintang + "*";
  }
  console.log(bintang);
}
