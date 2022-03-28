/**
 * Buatlah looping menggunakan "for" dan "while" dengan ketentuan berikut:
 * - Loop "for" dan "while" dilakukan 2 kali
 * - loop yang pertama dimulai dari 1 sampai 20 dan menampilkan "<index> - I love coding"
 * - loop yang kedua dimulai dari 20 sampai 1 dan menampilkan "<index> - I will become frontend developer"
 */

// 1. Melakukan Looping Menggunakan For
var hasil = "";

console.log("LOOPING FOR PERTAMA");
for(var i = 1; i <= 20; i++) {
  hasil=`${i} - I love coding `;
  console.log(hasil);
}

console.log("LOOPING FOR KEDUA");
for(var i = 20; i >= 1; i--) {
  hasil = `${i} - I will become frontend developer`;
  console.log(hasil);
}

console.log();

// 2. While
var result = "";

console.log("LOOPING WHILE PERTAMA");
var i=1;
while(i<=20) {
  result = `${i} - I love coding `;
  i++;
  console.log(result);
}

console.log("LOOPING WHILE KEDUA");
var i=20;
while(i>=1) {
  result = `${i} - I will become frontend developer `;
  i--;
  console.log(result);
}
