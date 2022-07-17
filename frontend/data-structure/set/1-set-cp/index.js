// Tulis program sebagai fungsi yang memeriksa apakah dua set dikatakan intersection atau tidak.
// Jika kedua set intersection, fungsi tersebut harus mengembalikan nilai intersection nya.
//
// Contoh 1
// Input: a = {"Java", "Python", "Javascript", "C ++", "C#"}, b = {"Swift", "Java", "Kotlin", "Python"}
// Output: {"Java", "Python"}
// Explanation: intersection dari a dan b adalah "Java" dan "Python"
//
// Contoh 2
// Input: a = {"Java", "Python"}, b = {"Swift"}
// Output: {}
// Explanation: tidak ada intersection dari a dan b


function intersection(setA, setB) {
    let result = new Set();
    return result; // TODO: replace this
}

let setA = new Set(["Java", "Python", "Javascript", "C ++", "C#"]);
let setB = new Set(["Java", "Python"]);
console.log(intersection(setA, setB)); // {"Java", "Python"}

setA = new Set("Java", "Python");
setB = new Set("Swift");
console.log(intersection(setA, setB)); // {}

module.exports = { intersection };