// contoh penggunaan sederhana dari callback adalah seperti ini

function addExclamation(str) {
  return `${str}!!!`;
}

function greeting(name) {
  return `Hello, ${name}`;
}

// ini adalah function yang menggunakan konsep callback

function playString(param, cb) {
  return cb(param);
}

console.log(playString(greeting("Joko"), addExclamation)); // Hello, Joko!!!

// atau bisa juga dibalik

console.log(playString(addExclamation("Joko"), greeting)); // Hello, Joko!!!

// Poin yang ingin disampaikan dari contoh diatas adalah bahwa callback bukanlah spesifik sebuah function
// namun lebih kearah konsep.
// Ya, callback memang BERUPA function, tapi pengertian tersebut kadang mengarah pada pertanyaan "fuction yang seperti apa?"
// Pada intinya callback adalah sebuah konsep yang menggunakan 1 function (katakanlah function A) sebagai parameter function B
// kemudian menggunakan function tersebut untuk mengolah lebih jauh data yang dihasilkan dari function B
