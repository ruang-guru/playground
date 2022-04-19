// Pencarian linier (atau pencarian berurutan) adalah metode untuk menemukan elemen dalam daftar.
// Ini secara berurutan memeriksa setiap elemen daftar sampai kecocokan ditemukan atau seluruh daftar telah dicari.

// Algoritma
// Mulai dari elemen paling kiri arr[] dan satu per satu bandingkan x dengan setiap elemen arr[]
// Jika x cocok dengan elemen, return 'x found at position i+1'.
// Jika x tidak cocok dengan salah satu elemen, return 'x not found!'.

// Contoh
// [1,3,12,23,34,56,62,98,99]
// x = 34

// Iteration 1:
// compare x with element at position 0
// NOT SAME

// Iteration 2:
// compare x with element at position 1
// NOT SAME

// Iteration 3:
// compare x with element at position 2
// NOT SAME

// Iteration 4:
// compare x with element at position 3
// NOT SAME

// Iteration 5:
// compare x with element at position 4
// SAME
// Element is at position 4

function linearSearch(arr, search) {
    result = `${search} not found!`;
    for (let i=0; i<arr.length; i++) {
        if (search == arr[i]) {
            result = `${search} found at position ${i+1}`;
            return result
        }
    }
    return result
}

let nums = [1, 3, 12, 23, 34, 56, 62, 98, 99];
console.log(linearSearch(nums, 34));
console.log(linearSearch(nums, 33));