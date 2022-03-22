// Masukan 3 bilangan bulat menggunakan if condition
const num1 = parseInt(prompt("Masukan bilangan pertama: "));
const num2 = parseInt(prompt("Masukan bilangan kedua: "));
const num3 = parseInt(prompt("Masukan bilangan ketiga: "));

let largest;

// beginanswer
// Pengecekan kondisi
if(num1 >= num2 && num1 >= num3) {
    largest = num1;
} else if (num2 >= num1 && num2 >= num3) {
    largest = num2;
} else {
    largest = num3;
}
// endanswer

// Menampilkan hasil
console.log("Bilangan terbesar adalah " + largest);