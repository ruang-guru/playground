// Staircase
// - Problem Solving (Basic)
// - Difficulty: Easy

/*
 * Complete the 'staircase' function below.
 *
 * The function accepts INTEGER n as parameter.
 */

// Full Problem: https://www.hackerrank.com/challenges/staircase/problem

function staircase(n) {
  // Write your code here
  //beginanswer
  var stair = [];
  var count = 0;
  for (var i = 0; i < n; i++) {
    var horiz = [];

    for (var j = 0; j < n; j++) {
      if (i == count && j >= n - 1 - count) {
        horiz[j] = '#';
      } else {
        horiz[j] = ' ';
      }
    }
    count++;
    stair.push(horiz.join(''));
  }
  // console.log(stair.join('\n'));
  return stair.join('\n')
  //endanswer
}

function main() {
  const n = 6

  let result = staircase(n);
  console.log(result)
}

main(); // execute program

module.exports = staircase