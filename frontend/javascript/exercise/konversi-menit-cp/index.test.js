const konversiMenit = require("./main.js")

describe("konversiMenit", () => {
  it("should convert minute into hour with format hh:mm", () => {
    let case1 = konversiMenit(10)
    let case2 = konversiMenit(65)
    let case3 = konversiMenit(700)

    expect(case1).toBe("0:10")
    expect(case2).toBe("1:05")
    expect(case3).toBe("11:40")
  })
})
