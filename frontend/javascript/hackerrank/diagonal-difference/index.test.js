const { expect, describe, it } = require('@jest/globals');
const diagonalDifference = require('./main')

describe("Diagonal Difference", () => {
  it("should return the absolute difference between the sums of the matrix's two diagonals as a single integer.", () => {
    let case1 = diagonalDifference([
      [11, 2, 4],
      [4, 5, 6],
      [10, 8, -12]
    ])

    expect(case1).toBe(15)
  })
})