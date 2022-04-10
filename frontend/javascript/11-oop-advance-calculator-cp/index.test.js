const Calculator = require('./index')

describe("Advance Calculator", () => {
  describe("add", () => {
    it("should return the total of addition", () => {
      const calculator = new Calculator(10)

      expect(calculator.add(10).operand).toBe(20)
    })
  })

  describe("subtract", () => {
    it("should return the total of subtraction", () => {
      const calculator = new Calculator(10)

      expect(calculator.subtract(5).operand).toBe(5)
    })
  })

  describe("multiply", () => {
    it("should return the total of multiplication", () => {
      const calculator = new Calculator(10)

      expect(calculator.multiply(10).operand).toBe(100)
    })
  })

  describe("divide", () => {
    it("should return the total of division", () => {
      const calculator = new Calculator(10)

      expect(calculator.divide(5).operand).toBe(2)
    })
  })

  describe("square", () => {
    it("should return the total of square calculation", () => {
      const calculator = new Calculator(10)

      expect(calculator.square(2).operand).toBe(100)
    })
  })

  describe("square root", () => {
    it("should return the total of square root", () => {
      const calculator = new Calculator(144)

      expect(calculator.squareRoot().operand).toBe(12)
    })
  })

  describe("add then subtract", () => {
    it("should return the total of addition followed by subtraction", () => {
      const calculator = new Calculator(10)

      expect(calculator.add(10).subtract(5).operand).toBe(15)
    })
  })

  describe("multiply then divide", () => {
    it("should return the total of multiplication followed by division", () => {
      const calculator = new Calculator(10)

      expect(calculator.multiply(10).divide(5).operand).toBe(20)
    })
  })

  describe("square then square root", () => {
    it("should return the total of square calculation followed bt square root", () => {
      const calculator = new Calculator(12)

      expect(calculator.square(2).squareRoot().operand).toBe(12)
    })
  })
})