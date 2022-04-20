// Tulis program sebagai fungsi yang mengembalikan union dari dua set.
// Jika kedua set union, fungsi tersebut harus mengembalikan hasilnya.
//
// Contoh 1
// Input: a = [1, 2, 3, 4, 5], b = [2, 3, 5, 7, 11]
// Output: [1, 2, 3, 4, 5, 7, 11]
// Explanation: intersection dari a dan b adalah 2, 3 dan 5
//
// Contoh 2
// Input: a = [1, 2, 3], b = [4, 5, 6]
// Output: [1, 2, 3, 4, 5, 6]
// Explanation: tidak ada intersection dari a dan b

function union(setA, setB) {
  const unionSet = new Set();
  for (const elem of setA) {
    unionSet.add(elem);
  }
  for (const elem of setB) {
    unionSet.add(elem);
  }
  return unionSet;
}

let set1 = new Set([1, 2, 3, 4, 5]);
let set2 = new Set([2, 3, 5, 7, 11]);

console.log(union(set1, set2)); // [1, 2, 3, 4, 5, 7, 11]