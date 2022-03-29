const rotation = require("./study-case-1.js")

describe("rotation", () => {
  it("should return correct matrix", () => {
    let sut = rotation([[1,2,3],[4,5,6],[7,8,9]])

    expect(sut).toStrictEqual([[7,4,1],[8,5,2],[9,6,3]])
  })

  it("should return another correct matrix", () => {
    let sut = rotation([[5,1,9,11],[2,4,8,10],[13,3,6,7],[15,14,12,16]])

    expect(sut).toStrictEqual([[15,13,2,5],[14,3,4,1],[12,6,8,9],[16,7,10,11]])
  })

  it("should return invalid data for undefined input", () => {
    let sut = rotation()

    expect(sut).toBe("invalid data")
  })

  it("should return invalid data for empty array", () => {
    let sut = rotation([])

    expect(sut).toBe("invalid data")
  })

  it("should return warning for a non 2d array", () => {
    let sut = rotation([1, 2, 3])

    expect(sut).toBe("input should be a 2d array")
  })
})
