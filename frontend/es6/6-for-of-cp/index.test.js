const convertElementToType = require(".");

describe("ES6: for-of looping 1 ", () => {
  it("Should return array of types", () => {
    expect(convertElementToType(["A", 3, [2]])).toEqual([
      "string",
      "number",
      "object",
    ]);
  });
});
