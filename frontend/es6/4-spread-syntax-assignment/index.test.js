const mergeTwoObjects = require(".");

describe("ES6: Spread Syntax 1", () => {
  it("Should return merged array", () => {
    expect(mergeTwoObjects({ a: 1, b: 2 }, { c: 3, d: 4 })).toEqual({
      a: 1,
      b: 2,
      c: 3,
      d: 4,
    });
  });
});
