const countRepetition = require(".");

describe("ES6: Functional programming filter string", () => {
  it("Should return repetitive counts number", () => {
    expect(countRepetition([1, 1, 1, 5, 5, 10, 9])).toEqual({
      1: 3,
      5: 2,
      10: 1,
      9: 1,
    });
  });
});
