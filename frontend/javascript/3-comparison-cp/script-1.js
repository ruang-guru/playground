// Pengecekan angka positif, negatif, atau 0 dengan nested if

// Masukan suatu angka
const number = parseInt(prompt("Enter a number: "));

// check if number is greater than 0
if (number > 0) {
    console.log("Angka Positif");
}

// check if number is 0
else if (number == 0) {
  console.log("Angka Negatif");
}

// if number is less than 0
else {
     console.log("Angka Negatif");
}
