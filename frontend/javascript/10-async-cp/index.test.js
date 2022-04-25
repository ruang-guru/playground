const { getDataPeopleById } = require("./index");

describe("ASync/await: Get starwars character summary", () => {
  it("Found character should return data", async () => {
    const result = await getDataPeopleById(1); 
    expect(result).toEqual("Luke Skywalker, memiliki tinggi 172cm dan lahir pada tahun 19BBY");
  });
});
