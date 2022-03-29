const terjemahKataSandi = require("./main.js")

describe("terjemahKataSandi", () => {
  it("should return messages that have been decrypted", () => {
    let case1 = terjemahKataSandi("&P%&+^S&^U&K+%N#1]#M&^+J^%%U#K #+^R&& +#2]#J+M#3#S^%%O&^R #")

    expect(case1).toBe('PASUKAN 1, MAJU KE AREA 2, JAM 3 SORE ')
  })
})