const format = require(".");

describe("ES6: Destructuring Object 2 Checkpoint", function () {
  it("Name should return same as format", function () {
    expect(format({ name: "John", email: "john@example.com" })).toBe("John - john@example.com - Unemployed");
  });
});
