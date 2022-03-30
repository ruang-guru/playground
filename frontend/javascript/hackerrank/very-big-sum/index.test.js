const { expect, describe, it } = require('@jest/globals');
const aVeryBigSum = require('./main')

describe("A Very Big Sum", () => {
  it('should return the sum of all array elements', () => {
    let case1 = aVeryBigSum([10000000010, 10000000020, 10000000030, 10000000040, 10000000050])

    expect(case1).toBe(50000000150)
  })
})