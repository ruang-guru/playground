const { expect, describe, it } = require('@jest/globals');
const { groupByAge, ageDistribution } = require('./index');

describe("groupByAge", () => {
    describe("when given an array of people", () => {
        it("should return a map of age groups to arrays of people", () => {
            let people = [
                { name: "Bob", age: 21 },
                { name: "Sam", age: 28 },
                { name: "Ann", age: 21 },
                { name: "Joe", age: 22 },
                { name: "Ben", age: 28 },
            ];
            let expected = new Map();
            expected.set(21, [{ name: "Bob", age: 21 }, { name: "Ann", age: 21 }]);
            expected.set(22, [{ name: "Joe", age: 22 }]);
            expected.set(28, [{ name: "Sam", age: 28 }, { name: "Ben", age: 28 }]);
            expect(groupByAge(people)).toEqual(expected);
        });
    })
})

describe("ageDistribution", () => {
    describe("when given an array of people", () => {
        it("should return a map of age groups to the number of people in that age group", () => {
            let people = [
                { name: "Bob", age: 21 },
                { name: "Sam", age: 28 },
                { name: "Ann", age: 21 },
                { name: "Joe", age: 22 },
                { name: "Ben", age: 28 },
            ];
            let expected = new Map();
            expected.set(21, 2);
            expected.set(22, 1);
            expected.set(28, 2);
            expect(ageDistribution(people)).toEqual(expected);
        })
    })
})