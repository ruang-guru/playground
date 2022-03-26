/* 
Komandan pasukan mengirimkan pesan instruksi berupa berisi kata sandi ke pasukan tentaranya. Hanya pasukan tentaranya saja yang bisa menerjemahkan kata sandi.
Untuk bisa menerjemahkan kata sandi tersebut harus mengikuti aturan dibawah ini. 
Jika di dalam Kata Sandi mengandung:
1. Karakter "&" maka karakter tersebut dihapus
2. Karakter "%" maka karakter tersebut dihapus
3. Karakter "^" maka karakter tersebut dihapus
4. Karakter "#" maka karakter tersebut diganti " " (spasi)
5. Karakter "]" maka karakter tersebut diganti ","
6. Karakter "+" maka karakter tersebut diganti "A"
7. Karakter " " (spasi) maka karakter tersebut diganti "E"
8  Karakter lainnya tidak diganti
# CONTOHNYA
INPUT:
kataSandi = "&P%&+^S&^U&K+%N#1]#M&^+J^%%U#K #+^R&& +#2]#J+M#3#S^%%O&^R #"
OUTPUT:
terjemahKataSandi = "PASUKAN 1, MAJU KE AREA 2, JAM 3 SORE"
Lengkapilah function dengan input kata sandi dan output terjemahannya
*/


function terjemahKataSandi(kataSandi) {
  // TODO: answer here
}


console.log(terjemahKataSandi("&P%&+^S&^U&K+%N#1]#M&^+J^%%U#K #+^R&& +#2]#J+M#3#S^%%O&^R #"));
// PASUKAN 1, MAJU KE AREA 2, JAM 3 SORE

module.exports = terjemahKataSandi