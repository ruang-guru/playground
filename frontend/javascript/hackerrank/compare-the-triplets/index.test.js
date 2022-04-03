const { expect, describe, it } = require('@jest/globals');
const compareTriplets = require('./main')

describe("Compare the Triplets", () => {
  it("should returns an array of each score from the 2 inputs", () => {
    let case1 = compareTriplets([17, 28, 30], [99, 16, 8])

    expect(case1).toEqual([2, 1])
  })
})