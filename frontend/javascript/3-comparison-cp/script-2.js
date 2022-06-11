// Masukan 3 bilangan bulat menggunakan if condition
const num1 = parseInt(prompt("Masukan bilangan pertama: "));
const num2 = parseInt(prompt("Masukan bilangan kedua: "));
const num3 = parseInt(prompt("Masukan bilangan ketiga: "));

let largest = 0

// TODO: answer here

// Menampilkan hasil

// ADITYA RIZQI ARDHANA - FE2273128 - FE4
if(num1 > num2 && num1 > num3 ) {
    largest = num1;
} else if(num2 > num1 && num2 > num3) {
    largest = num2;
} else if (num3 > num2 && num3 > num1) {
    largest = num3;
} else {
    largest = NaN;
}
// ADITYA RIZQI ARDHANA

alert("Bilangan terbesar adalah " + largest);
console.log("Bilangan terbesar adalah " + largest);