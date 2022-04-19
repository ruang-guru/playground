const constantNoChange = require("./index")

describe("ES6: Const", function () {
  it("Should return error", function () {
    expect(() => {
      constantNoChange();
    }).toThrow();
  });
});
