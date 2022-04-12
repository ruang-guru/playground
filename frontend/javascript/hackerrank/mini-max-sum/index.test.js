const { expect, describe, it } = require('@jest/globals');
const miniMaxSum = require('./main')

describe("Mini-Max Sum", () => {
  it("should return the respective minimum and maximum values as a single line of two space-separated long integers.", () => {
    let case1 = miniMaxSum([1, 3, 5, 7, 9])

    expect(case1).toBe('16 24')
  })
})