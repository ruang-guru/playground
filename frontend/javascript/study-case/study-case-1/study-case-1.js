// Diberikan matriks n x n 2D yang ada di bawah, putar matriks 90 derajat (searah jarum jam).
// Kita harus memutar matriks tersebut, yang berarti harus mengubah input matriks 2D secara langsung. 
// Jangan mengalokasikan matriks 2D lain dan melakukan rotasi, dan matriks harus berbentuk array 2 dimensi.
//
// Contoh 1
//
// 1 2 3     7 4 1
// 4 5 6 --> 8 5 2
// 7 8 9     9 6 3
//
// Input: matrix = [[1,2,3],[4,5,6],[7,8,9]]
// Output: [[7,4,1],[8,5,2],[9,6,3]]
//
// Contoh 2
//
//  5  1  9 11        15 13  2  5
//  2  4  8 10  --->  14  3  4  1
// 13  3  6  7  --->  12  6  8  9
// 15 14 12 16        16  7 10 11
//
// Input: matrix = [[5,1,9,11],[2,4,8,10],[13,3,6,7],[15,14,12,16]]
// Output: [[15,13,2,5],[14,3,4,1],[12,6,8,9],[16,7,10,11]]


function rotation(matrix) {
    // beginanswer
    if (matrix === undefined || matrix.length == 0) {
        return "invalid data"
    }

    if (matrix[0].constructor === Array) {
        swapTopRightToBottomLeft(matrix)
        swapRightToLeft(matrix)

        return matrix
    } else {
        return "input should be a 2d array"
    }
    // endanswer
};

// beginanswer
function swapTopRightToBottomLeft(matrix) {
    
    const n = matrix.length - 1
    for (let i = 0; i <= n; i++) {
        for (let j = i; j <= n; j++) {
            const temp = matrix[i][j]
            matrix[i][j] = matrix[j][i]
            matrix[j][i] = temp
        }
    }
}

function swapRightToLeft(matrix) {
    
    const n = matrix.length - 1
    for (let i = 0; i <= n; i++) {
        for (let j = 0; j <= n / 2; j++) {
            const temp = matrix[i][j]
            matrix[i][j] = matrix[i][n - j]
            matrix[i][n - j] = temp
        }
    }
}
// endanswer

module.exports = rotation