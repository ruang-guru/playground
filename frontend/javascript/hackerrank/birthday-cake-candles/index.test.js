const { expect, describe, it } = require('@jest/globals');
const birthdayCakeCandles = require('./main')

describe("Birthday Cake Candles", () => {
  it("should return the number of candles that are tallest", () => {
    let case1 = birthdayCakeCandles([ 28, 36, 28, 28])

    expect(case1).toBe(1)
  })
})