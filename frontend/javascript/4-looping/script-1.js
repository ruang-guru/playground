// Mengembalikan urutan bilangan, dan mengetahui apakah bilangan genap/ganjil

// Masukan jumlah bilangan yang ingin dicek, iterasi dari angka 1
const counter = parseInt(prompt("Masukan jumlah bilangan yang ingin dicek: "));

for (var x = 1; x <= counter; x++) {
    if (x % 2 === 0) {
        console.log(x + " adalah bilangan genap");   
    } else {
        console.log(x + " adalah bilangan ganjil");
    }
}