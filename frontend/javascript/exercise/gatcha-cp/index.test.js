const gatcha = require("./main.js")

describe("Gatcha", () => {
  it("should return to try again", () => {
    let case1 = gatcha(1)

    expect(case1).toBe('coba lagi ya')
  })

  it("should obtain 5 coupons", () => {
    let case2 = gatcha(2)

    expect(case2).toBe('selamat kamu mendapatkan kupon sebanyak 5')
  })

  it("should obtain 15 coupons", () => {
    let case3 = gatcha(3)

    expect(case3).toBe('selamat kamu mendapatkan kupon sebanyak 15')
  })

  it("should obtain 50 coupons", () => {
    let case4 = gatcha(4)

    expect(case4).toBe('selamat kamu mendapatkan kupon sebanyak 50')
  })

  it("should obtain 100 coupons", () => {
    let case5 = gatcha(5)

    expect(case5).toBe('selamat kamu mendapatkan kupon sebanyak 100')
  })

})
