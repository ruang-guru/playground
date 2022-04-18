const filterString = require(".");

describe("ES6: Functional programming filter number", () => {
  it("Should return element string", () => {
    expect(filterString(["s", "df", "g", 9, 10, "1D1", 27])).toEqual([
      "s",
      "df",
      "g",
      "1D1",
    ]);
  });
});
