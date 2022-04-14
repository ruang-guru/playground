const customisableGreeting = require("./index");

describe("ES6: Higher Order Function", function () {
  it("Name should return same as format", function () {
    expect(
      customisableGreeting(
        {
          firstName: "John",
          lastName: "Doe",
        },
        ({ firstName, lastName }) => {
          return `${firstName} ${lastName}`;
        }
      )
    ).toBe("Hi name is John Doe, how are you?");
  });
});


