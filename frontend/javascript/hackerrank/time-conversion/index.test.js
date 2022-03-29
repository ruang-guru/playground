const { expect, describe, it } = require('@jest/globals');
const timeConversion = require('./main')

describe("Time Conversion", () => {
  it("should return a new string representing the input time in 24 hour format.", () => {
    let case1 = timeConversion('12:20:50AM')
    let case2 = timeConversion('04:25:00PM')

    expect(case1).toBe('00:20:50')
    expect(case2).toBe('16:25:00')
  })
})