const max = require(".");

describe("ES6: Spread Syntax 2", () => {
  it("Should return max value", () => {
    expect(max(1)).toEqual(1);
    expect(max(1, 5, 6, 7)).toEqual(7);
    expect(max(1, 5, 6, 7, 1, 5, 6, 7, 1, 5, 6, 7, 29)).toEqual(29);
  });
});
