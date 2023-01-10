// Menukar variabel a dan b dengan operator matematika

let a = parseInt(prompt('Masukan variabel pertama: '));
let b = parseInt(prompt('Masukan variabel kedua: '));

// TODO: answer here
a = b + a;
b = a - b;
a = a - b;

console.log(`Nilai a setelah ditukar: ${a}`);
console.log(`Nilai b setelah ditukar: ${b}`);
