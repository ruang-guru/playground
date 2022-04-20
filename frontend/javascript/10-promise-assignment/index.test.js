const { getDataPeopleByIdWithFilms } = require("./index");

describe("Promise: Get starwars character with films ", () => {
  it("Found character should return data", () => {
    return getDataPeopleByIdWithFilms(1).then((data) => {
      expect(data.name).toEqual("Luke Skywalker");
      expect(data.films[0].title).toEqual("A New Hope");
      expect(data.films[data.films.length - 1].title).toEqual(
        "Revenge of the Sith"
      );
    });
  });
});
