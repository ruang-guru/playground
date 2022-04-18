const { getStarshipModelByCharacterId } = require("./index");

jest.useRealTimers();

describe("Async/await: Get starships by character id", () => {
  it("Found starships should return data", async () => {
    const result = await getStarshipModelByCharacterId(1); 
    expect(result).toEqual(["T-65 X-wing", "Lambda-class T-4a shuttle"]);
  });
});
