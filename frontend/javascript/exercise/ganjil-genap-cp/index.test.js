const ganjilGenap = require("./main.js")

describe("ganjilGenap", () => {
  it("should return both odd and even number", () => {
    let case1 = ganjilGenap('2341;3429;864;1309;1276')

    expect(case1).toBe('plat genap sebanyak 2 dan plat ganjil sebanyak 3')
  })

  it("should return only odd number", () => {
    let case2 = ganjilGenap('2349;3427;8645')

    expect(case2).toBe('plat ganjil sebanyak 3 dan plat genap tidak ditemukan')
  })

  it("should return only even number", () => {
    let case3 = ganjilGenap('2344;3428')

    expect(case3).toBe('plat genap sebanyak 2 dan plat ganjil tidak ditemukan')
  })

  it("should return not found", () => {
    let case4 = ganjilGenap('')

    expect(case4).toBe('plat tidak ditemukan')
  })

  it("should return invalid data", () => {
    let case5 = ganjilGenap()

    expect(case5).toBe('invalid data')
  })

})
