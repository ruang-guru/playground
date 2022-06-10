// Mengembalikan teks yang berkebalikan
//
// contoh:
// teks = hello world
// hasil:
// dlrow olleh

function reverseString(str) {
  // TODO: answer here
  var splitString = str.split("");
  var reverseArray = splitString.reverse();
  var joinArray = reverseArray.join("");

  return joinArray;
}

const string = prompt("Masukan teks: ");

const result = reverseString(string);
console.log(result);
