// Menggunakan switch case, petakan nilai ujian di bawah ini ke dalam Nilai A B C D E
// A = 91-100
// B = 71-90
// C = 61-70
// D = 51-60
// E = <= 50

// Masukan suatu angka
const score = parseInt(prompt('Masukan nilai: '));

// TODO: answer here
if (100 >= score > 90) {
  console.log('A');
} else if (91 > score > 70) {
  console.log('B');
} else if (71 > score > 50) {
  console.log('C');
} else if (51 > score > 30) {
  console.log('D');
} else if (50 >= score) {
  console.log('E');
}
