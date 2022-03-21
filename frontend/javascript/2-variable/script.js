// Menukar variabel a dan b

let a = prompt('Masukan variabel pertama: ');
let b = prompt('Masukan variabel kedua: ');

// 1. Buat variabel temp, variabel sementara untuk menuampung nilai yang ditukar
let temp;

// 2. Proses menukar variabel
temp = a;
a = b;
b = temp;

console.log(`Nilai a setelah ditukar: ${a}`);
console.log(`Nilai b setelah ditukar: ${b}`);