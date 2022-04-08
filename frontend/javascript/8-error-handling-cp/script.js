// Buatlah fungsi yang mengembalikan nama bulan dari array berdasarkan nilai bilangan yang dikirim ke fungsi tersebut. 
// Jika nilainya tidak sesuai dengan bulan (1-12), kembalikan error 'Masukan salah. Input bulan antara 1-12'.
//
// Contoh 1
// Input: 3
// Output: "Bulan ke-3 adalah Maret"
//
// Contoh 2
// Input: 15
// Output: "Masukan salah. Input bulan antara 1-12"

function getMonthName(monthNumber) {
    monthNumber = monthNumber - 1;
    let months = ['Januari', 'Februari', 'Maret', 'April', 'Mei', 'Juni', 'Juli',
                  'Agustus', 'September', 'Oktober', 'November', 'Desember'];
    if (months[monthNumber]) {
      console.log("Bulan ke-" + (monthNumber + 1) + " adalah " + months[monthNumber])
    } else {
      throw new Error("Masukan salah. Input bulan antara 1-12");
    }
  }
  
  // TODO: answer here
    let myMonth = parseInt(prompt("Masukan bilangan: "));
    getMonthName(myMonth);
  // TODO: answer here