// Pengecekan angka positif, negatif, atau 0 dengan nested if

// Masukan suatu angka
const number = parseInt(prompt('Masukan bilangan: '));

// TODO: answer here
if (number > 0) {
  console.log('bilangan ini adalah bilangan positive');
}

// check if number is 0
else if (number == 0) {
  console.log('bilangan ini adalah bilangan nol');
}

// if number is less than 0
else {
  console.log('bilangan ini adalah bilangan negative');
}
