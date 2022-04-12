const ruangCafe = require('./main.js')

describe("ruangCafe", () => {
  it("should return prohibited to enter", () => {
    let case1 = ruangCafe('',26,10000000)

    expect(case1).toBe('Anda tidak boleh masuk!')
  })

  it("should drink juice and display remaining money", () => {
    let case2 = ruangCafe('Karin',15,100000)

    expect(case2).toBe('Anda bisa pesan juice. Sisa uang anda: 50000')
  })

  it("should drink latte and display remaining money", () => {
    let case3 = ruangCafe('Naila',21,350000)

    expect(case3).toBe('Anda bisa pesan latte. Sisa uang anda: 50000')
  })

  it("should return not enough money", () => {
    let case4 = ruangCafe('Indah',9,40000)
    let case5 = ruangCafe('Inezka',30,20000)

    expect(case4).toBe('Uang tidak cukup. Anda harus pulang.')
    expect(case5).toBe('Uang tidak cukup. Anda harus pulang.')
  })

})