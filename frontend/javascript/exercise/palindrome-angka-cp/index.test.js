const angkaPalindrome = require('./main.js')

describe("angkaPlindrome", () => {
  it('should return to the nearest larger integer', () => {
    let case1 = angkaPalindrome(9)
    let case2 = angkaPalindrome(12)
    let case3 = angkaPalindrome(347)

    expect(case1).toBe(11)
    expect(case2).toBe(22)
    expect(case3).toBe(353)
  })
})