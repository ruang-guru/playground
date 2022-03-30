// Tulis kode untuk melakukan operasi akar dari suatu bilangan.
// Jika nilai yang diinput negatif, kembalikan pesan error

function squareRoot(number) {
    if(number < 0) {
        throw new Error("Maaf, operasi akar dengan bilangan negatif tidak dapat dilakukan");
    } else {
        return Math.sqrt(number);
    }
}
    
try {
    squareRoot(625);
    squareRoot(-9);
    squareRoot(100);
    
    console.log("Semua operasi akar berjalan dengan benar");
} catch(e) {
    console.log(e.message);
}