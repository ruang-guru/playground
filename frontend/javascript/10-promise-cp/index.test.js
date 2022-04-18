const {
  requestStarWarsPeopleById,
  promiseStarWarsPeopleById,
} = require("./index");

jest.useRealTimers();

describe("Promise: Get starwars character API ", () => {
  it("Non Promise: Found character should return data", () => {
    requestStarWarsPeopleById(
      1,
      (data) => {
        expect(data.name).toEqual("Luke Skywalker");
      },
      (e) => {}
    );
  });

  it("Found character should return data", () => {
    return promiseStarWarsPeopleById(1).then((data) => {
      expect(data.name).toEqual("Luke Skywalker");
    });
  });

  it("Not Found character should return error", () => {
    return promiseStarWarsPeopleById(999).catch((error) => {
      expect(error.toString()).toEqual("Error: 404");
    });
  });
});
