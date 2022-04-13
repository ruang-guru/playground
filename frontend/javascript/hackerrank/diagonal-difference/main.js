// Diagonal Difference
// - Problem Solving (Basic)
// - Difficulty: Easy

/*
* Complete the 'diagonalDifference' function below.
*
* The function is expected to return an INTEGER.
* The function accepts 2D_INTEGER_ARRAY arr as parameter.
*/

// Full Problem: https://www.hackerrank.com/challenges/diagonal-difference/problem

function diagonalDifference(arr) {
  // Write your code here
  //beginanswer
  let n = arr.length

  let sumDiagonal1 = 0;
  let sumDiagonal2 = 0;
  let cursorDiagonal1 = 0;
  let cursorDiagonal2 = n - 1;

  let count = 0;
  for (var i = 0; i < n; i++) {
    for (var j = 0; j < n; j++) {

      if (cursorDiagonal1 == count) {
        sumDiagonal1 += arr[i][j];
        cursorDiagonal1 = cursorDiagonal1 + n + 1;
      }

      if (cursorDiagonal2 == count && cursorDiagonal2 != (n * n) - 1) {

        sumDiagonal2 += arr[i][j];
        cursorDiagonal2 = cursorDiagonal2 + n - 1;
      }
      count++;
    }

  }

  let result = Math.abs(sumDiagonal1 - sumDiagonal2);
  return result
  //endanswer
}

function main() {
  //var n = parseInt(readLine());
  var a = [
    [11, 2, 4],
    [4, 5, 6],
    [10, 8, -12]
  ]; // override input
  let result = diagonalDifference(a);
  console.log(result);
}

main(); // execute program

module.exports = diagonalDifference