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

// beginanswer
let stars = "";
for (let i = 0; i < n; i++) {
  for (let k = 0; k < n - i; k++) {
    stars += "*";
  }
  stars += "\n";
}
console.log(stars);
// endanswer
