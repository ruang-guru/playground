/**
 * Tuliskan PSEUDOCODE Untuk menyelesaikan kasus berikut:
 *
 * Pada tahun 2020, Sebuah kebun binatang ingin mengganti harga tiketnya. Saat ini, kebun binatang tersebut memiliki HARGA DASAR
 * Rp 10.000. Harga tiket masuk akan disesuaikan dengan umur dari pengunjung tersebut. Kamu diminta untuk membuat program yang akan menghitung
 * harga tiket masuk dari tiap pengunjung. Di program ini nantinya, pengunjung akan menginput NAMA dan TAHUN KELAHIRAN.
 *
 * -Umur < 2 tahun: gratis
 * -Umur 2-10:  Harga dasar
 * -Umur 11-18:  Harga dasar dikalikan 1.5
 * -Umur 19 keatas: Harga dasar dikalikan 2
 * -Jika umurnya diatas 120 tahun ATAU dia kelahiran dibawah tahun 1900, maka tampilkan 'Invalid Age'
 *
 * Setelah menghitung harga, maka tampilkan NAMA dan HARGA TIKET dari pengunjung tersebut.
 * 
 *
 */

// TULIS PSEUDOCODE KAMU DIBAWAH INI:
/*
STORE "nama" with any value
STORE "tahun" with any value
STORE "harga" with 10000

IF "tahun" minus 2020 < 2
  DISPLAY "'name' gratis"
ELSE IF "tahun" minus 2020 >= 2 and <= 10
  DISPLAY "'name' with 'harga'"
ELSE IF "tahun" minus 2020 >= 11 and <= 18
  CALCULATE "harga" times 1.5
  DISPLAY "name" and calculation result 
ELSE IF "tahun" minus 2020 >= 19
  CALCULATE "harga" times 2
  DISPLAY "name" and calculation result
ELSE IF "tahun" minus 2020 > 120 or "tahun" <= 1900
  DISLPLAY "Invalid Age"
*/