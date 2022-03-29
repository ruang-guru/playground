const getMaximumGold = require("./study-case-2.js")

describe("get maximum gold", () => {

  it("should return correct golds", () => {
    let sut = getMaximumGold([[0,6,0],[5,8,7],[0,9,0]])

    expect(sut).toBe(24)
  })

  it("should return another correct golds", () => {
    let sut = getMaximumGold([[1,0,7],[2,0,6],[3,4,5],[0,3,0],[9,0,20]])

    expect(sut).toBe(28)
  })

  it("should return invalid data for undefined input", () => {
    let sut = getMaximumGold()

    expect(sut).toBe("invalid data")
  })

  it("should return invalid data for empty array", () => {
    let sut = getMaximumGold([])

    expect(sut).toBe("invalid data")
  })

  it("should return warning for a non 2d array", () => {
    let sut = getMaximumGold([1, 2, 3])

    expect(sut).toBe("input should be a 2d array")
  })
})
