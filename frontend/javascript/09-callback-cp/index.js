// Buatlah callback yang memenuhi function call di line 29

// function ini menduplikasi tiap karakter pada sebuah string
// contoh: hehe => hheehhee
function doubleChars(str) {
  // kerjakan disini
  //beginanswer
  return str
    .split("")
    .map((c) => `${c}${c}`)
    .join("");
  //endanswer
}

// function ini mengulang pengaplikasian callback pada str sejumlah num
function repeat(str, num, cb) {
  // kerjakan disini
  //beginanswer
  let result = str;

  for (i = 0; i < num; i++) {
    result = cb(result);
  }

  return result;
  //endanswer
}

console.log(repeat("triple", 2, doubleChars)); // ttttrrrriiiipppplllleeee
