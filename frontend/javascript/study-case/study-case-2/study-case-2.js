// Dalam suatu tambang emas ukuran m x n (2d array), setiap petak di tambang ini memiliki bilangan bulat yang mewakili jumlah emas di petak itu, 0 jika kosong.
// Kembalikan jumlah maksimum emas yang dapat dikumpulkan dengan ketentuan:
// 1. Setiap kali berada di petak, kamu akan mengumpulkan semua emas di petak itu.
// 2. Dari posisi sekarang, kamu bisa berjalan satu langkah ke kiri, kanan, atas, atau bawah.
// 3. Kamu tidak dapat mengunjungi petak yang sama lebih dari sekali.
// 4. Jangan pernah mengunjungi petak dengan 0 emas.
// 5. Kamu dapat memulai dan berhenti mengumpulkan emas dari posisi mana pun di petak yang memiliki emas.
//
// Contoh 1
// Input: grid = [[0,6,0],[5,8,7],[0,9,0]]
// Output: 24
// Explanation:
// 0 6 0
// 5 8 7
// 0 9 0
// Jalan untuk mendapatkan emas maksimal, 9 -> 8 -> 7.
//
// Contoh 2
// Input: grid = [[1,0,7],[2,0,6],[3,4,5],[0,3,0],[9,0,20]]
// Output: 28
// Explanation:
// 1  0  7
// 2  0  6
// 3  4  5
// 0  3  0
// 9  0 20
// Jalan untuk mendapatkan emas maksimal, 1 -> 2 -> 3 -> 4 -> 5 -> 6 -> 7.

function getMaximumGold(grid) {
    // beginanswer

    if (grid === undefined || grid.length == 0) {
        return "invalid data"
    }

    if (grid[0].constructor === Array) {
        const m = grid.length;
        const n = grid[0].length;
        let max = 0;
    
        for (let i = 0; i < m; i++) {
            for (let j = 0; j < n; j++) {
                if (grid[i][j] != 0) {
                    const sum = backtrack(grid, i, j, m, n);
                    max = Math.max(sum, max);
                }
            }
        }
        
        return max;
    } else {
        return "input should be a 2d array"
    }
    // endanswer
}; 

// beginanswer
function backtrack(grid, row, col, m, n) {
    if (outOfBound(row, col, m, n) || grid[row][col] === 0) return 0;
    
    let sum = grid[row][col];
    grid[row][col] = 0;  // penanda sebagai yang sudah dikunjungi
    
    const top = backtrack(grid, row - 1, col, m, n);
    const right = backtrack(grid, row, col + 1, m, n);
    const bot = backtrack(grid, row + 1, col, m, n);
    const left = backtrack(grid, row, col - 1, m, n);
    
    grid[row][col] = sum;
    
    return sum + Math.max(top, right, bot, left);
}

function outOfBound(row, col, m, n) {
    return row < 0 || col < 0 || row >= m || col >= n;
}
// endanswer

module.exports = getMaximumGold