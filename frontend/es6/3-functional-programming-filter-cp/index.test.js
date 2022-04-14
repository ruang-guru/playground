const findAnimal = require("./index")

describe("ES6: Functional programming filter string", () => {
  it("Should return only lowercase as hidden words", () => {
    expect(findAnimal("UkUNFYGaFYFYmNUHbJKHJJiOKDJKDnKFKFLgLLF")).toBe(
      "kambing"
    );
  });
});