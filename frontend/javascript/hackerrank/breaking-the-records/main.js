// Breaking the Records
// - Problem Solving (Basic)
// - Difficulty: Easy

/*
 * Complete the 'breakingRecords' function below.
 *
 * The function is expected to return an INTEGER_ARRAY.
 * The function accepts INTEGER_ARRAY scores as parameter.
 */


// Full Problem: https://www.hackerrank.com/challenges/breaking-best-and-worst-records/problem

function breakingRecords(scores) {
  // Write your code here
  //beginanswer
  let highestScore = scores[0];
  let lowestScore = scores[0];
  let arr = [0, 0];
  scores.forEach(score => {
      if (highestScore < score) {
          highestScore = score;
          arr[0] = arr[0]+1;
      }
      if (lowestScore > score) {
          lowestScore = score;
          arr[1] = arr[1]+1;
      }
  });

  return arr
  //endanswer
}

function main() {
  //     var n = parseInt(readLine());
  //     s = readLine().split(' ');
  //     s = s.map(Number);
  var scores = [10, 5, 20, 20, 4, 5, 2, 25, 1]; // override input
  var result = breakingRecords(scores);
  console.log(result.join(' ') + '\n')

}

main(); // execute program

module.exports = breakingRecords