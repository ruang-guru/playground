const checkOmnipresent = require(".");

describe("ES6: for-of looping 1 ", () => {
  it("Check if a value is omnipresen", () => {
    expect(
      checkOmnipresent(
        [
          [1, 2, 3],
          [5, 4, 1],
          [4, 14, 1],
        ],
        1
      )
    ).toEqual(true);
    expect(
      checkOmnipresent(
        [
          [1, 2, 3],
          [5, 4, 1],
          [4, 14, 1],
        ],
        2
      )
    ).toEqual(false);
  });
});
