const { multiplyByTwo, multiply, returnTheTwo} = require("./index")

describe("ES6: Arrow functio checkpoint", function () {
  it("All arrow function should return the correct value", function () {
    let value = 2;

    value = multiplyByTwo(value);
    value = multiply(value, value);
    value = returnTheTwo() + value;

    expect(value).toEqual(18);
  });
});