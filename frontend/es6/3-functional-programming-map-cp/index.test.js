const lowerToUpperMap = require("./index");

describe("ES6: Functional programming filter number", () => {
  it("Should should return filtered digit", () => {
    expect(lowerToUpperMap(["a", "b", "c", "d"])).toEqual([
      { a: "A" },
      { b: "B" },
      { c: "C" },
      { d: "D" },
    ]);
  });
});