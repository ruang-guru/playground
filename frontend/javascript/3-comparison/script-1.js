// Pengecekan bilangan positif, negatif, atau 0 dengan if..else if..else

// Masukan suatu bilangan
const number = parseInt(prompt("Masukan bilangan: "));

// Pengecekan jika bilangan lebih besar dari 0
if (number > 0) {
  console.log("bilangan positif");
}

// Pengecekan jika bilangan sama dengan 0
else if (number == 0) {
  console.log("bilangan = 0");
}

// Pengecekan jika bilangan kurang dari 0
else {
  console.log("bilangan negatif");
}