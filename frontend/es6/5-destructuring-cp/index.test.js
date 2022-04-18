const format = require("./index");

describe("ES6: Destructuring Object 1 Checkpoint", function () {
  it("Name should return  the same as format", function () {
    expect(format({ name: "John", email: "john@example.com" })).toBe("John - john@example.com");
  });
});
