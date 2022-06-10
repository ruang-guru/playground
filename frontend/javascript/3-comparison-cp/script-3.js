// Menggunakan switch case, petakan nilai ujian di bawah ini ke dalam Nilai A B C D E
// A = 91-100
// B = 71-90
// C = 61-70
// D = 51-60
// E = <= 50

// Masukan suatu angka
const score = parseInt(prompt("Masukan nilai: "));
console.log(score);

if (score >= 91 && score <= 100) {
  console.log("Nilai A");
}
// TODO: answer here
switch (score) {
  case score >= 91 && score <= 100:
    console.log("Nilai A");
    break;
  case score >= 71 && score <= 90:
    console.log("Nilai B");
    break;
  case score >= 61 && score <= 70:
    console.log("Nilai C");
    break;
  case score >= 51 && score <= 60:
    console.log("Nilai D");
    break;
  default:
    console.log("Nilai E");
    break;
}
