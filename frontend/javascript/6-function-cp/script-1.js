// Mengembalikan teks yang berkebalikan
//
// contoh: 
// teks = hello world
// hasil:
// dlrow olleh

const prompt = require("prompt-sync")({sigint:true});
function reverseString(str) {
var newString = "";
    for (var i = str.length - 1; i >= 0; i--) {
        newString += str[i];
    }
    return newString;
}

const string = prompt('Masukan teks: ');

const result = reverseString(string);
console.log(result);
