/*
Buat lah Algoritma, Peudocode dan Code untuk program Kampus Merdeka.
Program tersebut akan membatu untuk melihat apakah mahasiswa tersebut lulus atau tidak.
Setiap data terdapat nama, nilai, dan jumlah absen.
ketentuan lulus atau tidaknyanya mahasiswa sebagai berikut:
- Nilai 70 keatas
- Absensi dibawah 5
Jika mahasiswa dinyatakan lulus maka akan menampilkan "<nama> lulus"
Jika mahasiswa dinyatakan tidak lulus maka akan menampilkan "<nama> tidak lulus"
*/

// ALGORITMA:

// 1. Buat variabel "nama" dan masukan nama murid
// 2. Buat variabel "nilai" dan masukan nilai murid
// 3. Buat variabel "absen" dan masukan absen murid
// 4. Jika nilai 70 keatas dan absen dibawah 5 maka murid dinyatakan 'lulus'
// 5. Dan jika diluar dari kondisi tersebut murid dinyatakan 'tidak lulus'

// PSEUDOCODE:

// STORE "nama" with any value
// STORE "nilai" with any value
// STORE "absen" with any value
// IF "nilai" >= 70 and "absen" < 5
// DISPLAY "nama" lulus
// ELSE
// DISPLAY "nama" tidak lulus

var nama = "dito";
var nilai = 60;
var absen = 4;

if (nilai >= 70 && absen < 5) {
  console.log(`${nama} lulus`)
}
else {
  console.log(`${nama} tidak lulus`)
}