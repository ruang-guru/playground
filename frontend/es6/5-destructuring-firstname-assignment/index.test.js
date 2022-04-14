const getUserFirstName = require(".");

describe("ES6: Destructuring Object 3 Checkpoint", function () {
  it("User first name should be John", function () {
    expect(getUserFirstName({ name: "John", email: "john@example.com" })).toBe("John");
  });
});
