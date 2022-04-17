const returnUserAverageScore = require(".");

describe("ES6: Functional programming filter object", () => {
  it("Should return average score", () => {
    expect(
      returnUserAverageScore([
        {
          name: "Yono",
          nilai: [6, 8, 7, 7, 5, 5],
        },
        {
          name: "Dodi",
          nilai: [9, 6, 5, 10, 8, 6],
        },
        {
          name: "Dori",
          nilai: [6, 8, 7, 3, 5, 5],
        },
      ])
    ).toEqual([
      { name: "Yono", rataNilai: 6 },
      { name: "Dodi", rataNilai: 7 },
      { name: "Dori", rataNilai: 6 },
    ]);
  });
});
