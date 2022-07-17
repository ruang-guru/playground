const { expect, describe, it } = require('@jest/globals');
const breakingRecords = require('./main')

describe("Breaking the Records", () => {
  it("should return an array with the number of times the record was broken", () => {
    let case1 = breakingRecords([3, 4, 21, 36, 10, 28, 35, 5, 24, 42])

    expect(case1).toEqual([4, 0])
  })
})