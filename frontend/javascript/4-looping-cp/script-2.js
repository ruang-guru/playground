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
// const n = parseInt(prompt("Masukan jumlah baris: "));
const n = 5;
// TODO: answer here
for (let i = 1; i <= n; i++) {
  let result = "";
  for (let j = i; j <= n; j++) {
    result += "*";
  }
  console.log(result);
}
