// Menghapus n elemen terakhir dari array
// contoh: 
// [1, 2, 3, 4, 5, 6, 7, 8, 9]
// jumlah yang dihapus = 3
// hasil:
// [1, 2, 3, 4, 5, 6]


let numInput = prompt("Masukan urutan bilangan dengan pemisah spasi:");
let numArray = numInput.split(" ");

let counter = parseInt(prompt("Masukan banyaknya bilangan terakhir yang ingin dihapus: "));

while (counter--) {
    numArray.pop();
}
  
console.log("Hasil:", arr);