// Diberikan bilangan bulat n, kembalikan true jika pangkat dua. Jika tidak, kembalikan false.
// Suatu bilangan bulat n adalah pangkat dua, jika terdapat bilangan bulat x sedemikian sehingga n == 2^x.
//
// Contoh 1:
// Input: n = 1
// Output: true
// Explanation: 2^0 = 1
//
// Contoh 2:
//
// Input: n = 16
// Output: true
// Explanation: 2^4 = 16
//
// Example 3:
// Input: n = 3
// Output: false

let isPowerOfTwo = function(n) {
    if (n === 1) {
        return true;
    } else if (n <= 0) {
        return false;
    }
    
    if (n % 2 !== 0) {
        return false;
    } else {
        return isPowerOfTwo(Math.floor(n/2));    
    }
};