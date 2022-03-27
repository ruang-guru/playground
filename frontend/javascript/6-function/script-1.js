// Mengembalikan nilai yang lebih kecil antara dua bilangan
//
// contoh: 
// min(2, 5) == 2
// min(3, -1) == -1
// min(1, 1) == 1

// alternatif 1
function min(a, b) {
  if (a < b) {
    return a;
  } else {
    return b;
  }
}

// alternatif 2
function min(a, b) {
  return a < b ? a : b;
}