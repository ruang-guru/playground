// Mengembalikan boolean untuk mengecek apakah teks palindrom
//
// contoh:
// teks1 = hello
// teks2 = madam
// teks3 = kasur ini rusak
// hasil:
// teks1 dibalik menjadi olleh, maka balikan akan false
// teks2 dibalik sama menjadi madam, maka balikan akan true
// teks3 dibalik sama menjadi kasur ini rusak, maka balikan akan true

function checkPalindrome(str) {
  // TODO: answer here
  var splitString = str.split("");
  var reverseArray = splitString.reverse();
  var joinArray = reverseArray.join("");

  if (joinArray === str) {
    console.log("true");
  } else {
    console.log("false");
  }
  console.log(joinArray, "===", str);
}

// masukan teks
const string = prompt("Masukan string: ");

// memanggil fungsi palindrom
const value = checkPalindrome(string);

console.log(value);
