/*  
  Seseorang akan mendaftar untuk menjadi murid baru perguruan tinggi negeri Universitas di Indonesia. Untuk masuk universitas ini,
  harus mendapatkan nilai SNMPTN 80 dan keatas. Di dalam kampus ini juga terdapat 2 macam program:
  Nasional & Internasional. Program Nasional ditujukan untuk pendaftar berkewarganegaraan Indonesia dan program Internasional
  untuk kewarganegaraan lainnya.
  Para calon pendaftar ini akan menginput informasi sebagai berikut:
  Nama: Nama sesuai kartu Identitas.
  Citizenship: Kewarganegaraan, bisa 'Indonesia' atau negara lainnya.
  Nilai SNMPTN: nilai SNMPTN para pendaftar, dari skala 0-100.
  Program ini berjalan sebagai berikut:
  1. Jika nama kosong, tampilkan 'NAMA ANDA KOSONG!'TOLONG DICEK KEMBALI'
  2. Jika citizenship kosong, tampilkan 'TOLONG ISI KEWARGANEGARAAN ANDA SESUAI KARTU IDENTITAS!'
  3. Jika nilai SNMPTN kosong, maka tampilkan 'Nilai SNMPTN Kosong! ANDA TIDAK DAPAT MENGIKUTI PROSES SELEKSI'
  4. Program ini akan mengenerate hasil seleksi mahasiswa dengan format:
     -'<NAMA> LULUS SELEKSI PROGRAM NASIONAL' Jika pendaftar lulus seleksi & berkewarganegaraan 'Indonesia'.
     -'<NAMA> LULUS SELEKSI PROGRAM INTERNASIONAL' Jika pendaftar lulus seleksi & bukan warga negara 'Indonesia'.
     -'<NAMA> TIDAK LULUS SELEKSI' Jika pendaftar tidak lulus seleksi.
*/

 let name = 'revan'; // silakan berikan nilai bebas
 let citizenship = 'Sitka'; //silahkan berikan nilai bebas
 let grade = 79; //silahkan berikan nilai angka bebas;
 
 // TULIS KODE KAMU DIBAWAH INI:
 if(name === '') {
   console.log('NAMA ANDA KOSONG!TOLONG DICEK KEMBALI');
 }
 else if(citizenship === '') {
   console.log('TOLONG ISI KEWARGANEGARAAN ANDA SESUAI KARTU IDENTITAS!')
 }
 else if(grade === '') {
   console.log('Nilai SNMPTN Kosong! ANDA TIDAK DAPAT MENGIKUTI PROSES SELEKSI')
 }
 else if(citizenship === 'Indonesia' && grade >= 80) {
   console.log(`${name} LULUS SELEKSI PROGRAM NASIONAL`);
 }
 else if(citizenship !== 'Indonesia' && grade >= 80) {
   console.log(`${name} LULUS SELEKSI PROGRAM INTERNASIONAL`);
 }
 else if(grade < 80) {
   console.log(`${name} TIDAK LULUS SELEKSI`);
 }