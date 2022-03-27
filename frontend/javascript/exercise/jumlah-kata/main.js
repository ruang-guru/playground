// Cari lah jumlah kata yang terdapat pada kalimat berikut

function hitungJumlahKata(kalimat) {
  var hasil = "";
  var jumlah = 0;
  if (kalimat != "" && kalimat != " ") {
    jumlah = jumlah + 1;
  }
  for (i = 0; i < kalimat.length; i++) {
    if (kalimat[i] != " ") {
      hasil += kalimat[i];
    }
    else if (kalimat[i] === " " && hasil != "") {
      jumlah = jumlah + 1
    }
  }
  return jumlah
}

console.log(hitungJumlahKata('I have a dream')); // 4
console.log(hitungJumlahKata('I believe I can code')); // 5
console.log(hitungJumlahKata('I')); // 1