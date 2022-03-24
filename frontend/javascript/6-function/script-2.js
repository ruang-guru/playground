// Mengembalikan nilai pangkat n dari suatu bilangan x
//
// contoh: 
// pow(3, 2) = 3 * 3 = 9
// pow(3, 3) = 3 * 3 * 3 = 27
// pow(1, 100) = 1 * 1 * ...* 1 = 1

function pow(x, n) {
  let result = x;
  
  for (let i = 1; i < n; i++) {
    result *= x;
  }
  
  return result;
}
  
let x = parseInt(prompt("Masukan nilai x: "));
let n = parseInt(prompt("Masukan nilai n: "));
  
if (n < 1) {
  console.log("Pangkat tidak bisa dihitung, masukan bilangan bulat positif pada n");
} else {
  console.log(pow(x, n) );
}