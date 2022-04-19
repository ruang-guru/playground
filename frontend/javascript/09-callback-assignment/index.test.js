const { groupBy, isOdd } = require("./index")

describe("group-by", () => {
  it("should return groups of even and odd birth years", () => {
    let input = [
      {
        name: 'Dito',
        year: 1992,
        place: 'Bogor',
      },
      {
        name: 'Tegar',
        year: 1995,
        place: 'Depok',
      },
      {
        name: 'Uli',
        year: 1996,
        place: 'Bogor',
      },
    ];

    expect(groupBy(input, isOdd)).toEqual({ false:
      [ { name: 'Dito', year: 1992, place: 'Bogor' },
        { name: 'Uli', year: 1996, place: 'Bogor' } ],
     true: [ { name: 'Tegar', year: 1995, place: 'Depok' } ] })
  })
})