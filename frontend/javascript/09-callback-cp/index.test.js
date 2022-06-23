const { doubleChars, repeat } = require("./index")

describe("group-by", () => {
  it("should return a character string that has already been duplicated", () => {
    expect(repeat("triple", 2, doubleChars)).toMatch("ttttrrrriiiipppplllleeee")
  })
})