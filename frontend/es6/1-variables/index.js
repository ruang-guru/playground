"use strict";

// 1. Var merupakan function scope
function varTest(x) {
  // variabel scope di deklarasikan `var`
  var scope = "original";

  if (true) {
    // ketika ada variabel lain dengan nama sama dideklarasikan, maka nilai dari variabel akan ditimpa
    var scope = "overriden";
  }

  return scope; // overriden
}

// 2. Let merupakan block scope
function letTest() {
  // Variabel awal dideklarasikan dengan `let`
  let scope = "original";

  if (true) {
    // karena let block scope, maka pada tahap tidak mempengaruhi variable diluar scope
    let scope = "overriden";
    console.log(scope); // 'overriden'
  }

  // Keluaran fungsi ini akan tetap bernilai "original".
  return scope; // original
}

// 3. const tidak dapan di-reassign atau di override

function constTestOverride() {
  const scope = "original";

  if (true) {
    // Program akan berhenti dan mengeluarkan Refference error.
    scope = "overriden";
  }

  return scope;
}

// 4. block scope pada for loop

function forLoopBlockScope() {
  const list = [1, 2, 3, 4, 5];
  const arrayFunction = [];

  // Penggunaan var disini yang mana adalah function scope.
  for (var i = 0; i < list.length; i++) {
    arrayFunction.push(() => {
      return list[i];
    });
  }

  return arrayFunction[1](); //
  /**
   * Keluaran dari fungsi ini adalah 'undefined'.
   * Ketika fungsi didalam array di eksekusi menyebabkan nilai yang salah.
   */
}
