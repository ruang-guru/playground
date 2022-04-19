// Mengecek jika dua string adalah anagram
// Anagram adalah kata yang dibentuk melalui penataan ulang huruf-huruf dari beberapa kata lain.
//
// Contoh 1
// Input: a = "keen", b = "knee"
// Output: "Anagram"
// Explanation: Jika ditata, "knee" dan "keen" memiliki huruf-huruf yang sama, hanya berbeda urutan
//
// Contoh 2
// Input: a = "fried", b = "fired"
// Output: "Anagram"
// Explanation: Jika ditata, "fried" dan "fired" memiliki huruf-huruf yang sama, hanya berbeda urutan
//
// Contoh 3
// Input: a = "apple", b = "paddle"
// Output: "Bukan Anagram"
// Explanation: Jika ditata, "apple" dan "paddle" memiliki huruf-huruf yang berbeda

function anagramChecker(str1, str2) {
    return true; // TODO: replace this
}

console.log(anagramChecker("keen", "knee"));
console.log(anagramChecker("kee", "knn"));
console.log(anagramChecker("fried", "fired"));

module.exports = {
    anagramChecker
}
