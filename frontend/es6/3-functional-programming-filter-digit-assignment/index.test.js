const filterDigit = require("./index");

describe("ES6: Functional programming filter number", () => {
  it("Should return filtered digit", () => {
    expect(filterDigit([88, 44, 3, 8481, 444], 2)).toEqual([88, 44]);
  });
});
