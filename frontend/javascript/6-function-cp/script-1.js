// Mengembalikan teks yang berkebalikan
//
// contoh:
// teks = hello world
// hasil:
// dlrow olleh

function reverseString(str) {
  // TODO: answer here
  ubahString = str.toLowerCase();
  var pisahString = ubahString.split('');
  var balikArray = pisahString.reverse();
  var gabungArray = balikArray.join('');
  return gabungArray;
}

const string = prompt('Masukan teks: ');

const result = reverseString(string);
console.log(result);
