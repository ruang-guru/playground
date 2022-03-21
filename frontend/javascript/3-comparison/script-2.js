// Pengecekan angka positif, negatif, atau 0 dengan switch case

// Masukan suatu angka
const number = parseInt(prompt("Masukan bilangan: "));

// begin answer
switch (number) {
    case (number < 0):
        console.log("bilangan negatif");
        break;
    case (number == 0):
        console.log("bilangan = 0");
        break;
    case (number > 0):
        console.log("bilangan positif");
        break;
    default:
        break;
}
// end answer