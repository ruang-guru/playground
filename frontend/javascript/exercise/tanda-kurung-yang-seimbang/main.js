// 1. Tanda kurung dianggap sebagai salah satu dari karakter berikut: (, ), {, }, [, atau ].
// 2. Kita menganggap dua tanda kurung cocok jika elemen pertama adalah tanda kurung buka, misalnya, (, {, atau [, dan tanda kurung kedua adalah tanda kurung tutup dari jenis yang sama, misalnya, ( dan ), [ dan ], dan { dan } adalah satu-satunya pasangan tanda kurung yang cocok.
// 3. Selanjutnya, barisan kurung dikatakan seimbang jika memenuhi kondisi berikut: 
// - Barisan kosong, atau 
// - Barisan terdiri dari dua, tidak kosong, barisan keduanya seimbang, atau 
// - Kurung pertama dan terakhir dari urutannya cocok, dan porsi urutan diantaranya juga seimbang.
// 4. Kita diberi serangkaian tanda kurung. 
// 5. Tugas kita adalah menentukan apakah setiap urutan kurung seimbang. Jika string seimbang, kembalikan true, jika tidak, kembalikan false.
//
// Contoh 1
// s = {[()]}
// result: true
//
// Contoh 2
// s = {}()
// result: false
//
// Contoh 3
// s = {(})
// result: false
//
// Contoh 4
// s = )
// result: false


function isBalanced(s) {

  // Tulis kode di sini
  // beginanswer
  let stack = []; // O(n) space
  let OPENINGS = "{[(";
  let CLOSINGS = "}])";
  let PAIR = { ")": "(", "}": "{", "]": "[" };
  
  for (const letter of s) {
    // O(n)
    if (isOpening(letter)) {
      // O(n)
      pushInStack(letter); // O(1)
    } else if (isClosing(letter)) {
      // O(n)
      if (isOpeningPairFound(letter)) {
        stack.pop();
      } else {
        return false;
      }
    }
  }
  
  return stack.length < 1;
  
  function isOpening(letter) {
    return OPENINGS.indexOf(letter) > -1;
  }
  
  function pushInStack(letter) {
    stack.push(letter);
  }
  
  function isClosing(letter) {
    return CLOSINGS.indexOf(letter) > -1;
  }
  
  function isOpeningPairFound(closing) {
    if (stack.length < 1) return false;
    const opening = PAIR[closing];
    return opening === stack[stack.length - 1];
  }
  // endanswer
}
  