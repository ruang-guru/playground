// Bilangan Fibonacci, biasanya dilambangkan dengan F(n) membentuk suatu barisan, yang disebut barisan Fibonacci, 
// sehingga setiap bilangan adalah jumlah dari dua bilangan sebelumnya, dimulai dari 0 dan 1.
// F(0) = 0, F(1) = 1
// F(n) = F(n - 1) + F(n - 2), untuk n > 1.
// Diberikan n, hitung F(n).
//
// Contoh 1:
// Input: n = 2
// Output: 1
// Explanation: F(2) = F(1) + F(0) = 1 + 0 = 1.
//
// Contoh 2:
//
// Input: n = 3
// Output: 2
// Explanation: F(3) = F(2) + F(1) = 1 + 1 = 2.
//
// Example 3:
// Input: n = 4
// Output: 3
// Explanation: F(4) = F(3) + F(2) = 2 + 1 = 3.

let fib = function(n) {
    if(n === 0) return 0;
    if(n === 1) return 1;
    return fib(n-1) + fib(n-2);
};