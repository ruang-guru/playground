const sumTotalArray = require("./index");

describe("ES6: Functional programming filter string", () => {
  it("Should return sum of array", () => {
    expect(sumTotalArray([1, 1, 1, 5])).toBe(8);
  });
});