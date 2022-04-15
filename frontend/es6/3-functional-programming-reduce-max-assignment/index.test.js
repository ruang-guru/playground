const maxFromArray = require(".");

describe("ES6: Functional programming filter string", () => {
  it("Should return max of array", () => {
    expect(maxFromArray([1, 8, 1, 5, 3, 15])).toBe(15);
    expect(maxFromArray([1, 8, 100, 5, 3, 15])).toBe(100);
  });
});
