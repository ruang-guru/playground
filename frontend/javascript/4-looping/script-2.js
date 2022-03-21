// Mengembalikan pola * yang menambah di setiap barisnya
//
// contoh: 
// baris = 5
// hasil:
// *
// **
// ***
// ****
// *****


// Masukan jumlah baris
const n = parseInt(prompt("Masukan jumlah baris: "));

let stars = "";
for (let i = 1; i <= n; i++) {
  for (let j = 0; j < i; j++) {
    stars += "*";
  }
  stars += "\n";
}
console.log(stars);