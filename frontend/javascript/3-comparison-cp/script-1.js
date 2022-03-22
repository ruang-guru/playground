// Pengecekan angka positif, negatif, atau 0 dengan nested if

// Masukan suatu angka
const number = parseInt(prompt("Masukan bilangan: "));

// beginanswer
if (number >= 0) {
    if (number == 0) {
        console.log("bilangan = 0");
    } else {
        console.log("bilangan positif");
    }
} else {
    console.log("bilangan negatif");
}
// endanswer