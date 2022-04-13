/**
 * Kali ini kamu diminta untuk membuat sebuah game gatcha dengan ketentuan berikut:
 * - diberikan sebuah variable button yang akan mengenerate angka random 1 - 5
 * - jika mendapat angaka 1 maka akan menampilkan "coba lagi ya"
 * - jika mendapat angaka 2 maka akan menampilkan "selamat kamu mendapatkan kupon sebanyak 5"
 * - jika mendapat angaka 3 maka akan menampilkan "selamat kamu mendapatkan kupon sebanyak 15"
 * - jika mendapat angaka 4 maka akan menampilkan "selamat kamu mendapatkan kupon sebanyak 50"
 * - jika mendapat angaka 5 maka akan menampilkan "selamat kamu mendapatkan kupon sebanyak 100"
 * Buatlah Pseudocode beserta codenya menggunakan conditional "switch"
 */

// PSEUDOCODE:
//beginanswer
// SET 'button' with random value
// IF 'button' > 5
//  DO 'button' -5
// 
// SWITCH 'button'
//  case 1
//  DISPLAY "coba lagi ya"
//  case 2
//  DISPLAY "selamat kamu mendapatkan kupon sebanyak 5"
//  case 3
//  DISPLAY "selamat kamu mendapatkan kupon sebanyak 15"
//  case 4
//  DISPLAY "selamat kamu mendapatkan kupon sebanyak 50"
//  case 5
//  DISPLAY "selamat kamu mendapatkan kupon sebanyak 100"
//endanswer

function gatcha(button) {
  //beginanswer
  if (button > 5) {
    button = button - 5;
  }

  switch (button) {
    case 1: {
      return "coba lagi ya"
      break;
    }
    case 2: {
      return "selamat kamu mendapatkan kupon sebanyak 5"
      break;
    }
    case 3: {
      return "selamat kamu mendapatkan kupon sebanyak 15"
      break;
    }
    case 4: {
      return "selamat kamu mendapatkan kupon sebanyak 50"
      break;
    }
    case 5: {
      return "selamat kamu mendapatkan kupon sebanyak 100"
      break;
    }
  }
  //endanswer
}

// Create variable button here
//beginanswer
var button = Math.round(Math.random() * 10);
//endanswer

console.log(gatcha(button))

module.exports = gatcha
