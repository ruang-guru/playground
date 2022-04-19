const { expect, describe, it } = require('@jest/globals');
const { searchMatch } = require('./index');

describe("searchMatch", () => {
    describe("when two input arrays are empty", () => {
        it("returns an empty array", () => {
            let arr1 = [];
            let arr2 = [];
            let result = searchMatch(arr1, arr2);
            expect(result).toEqual([]);
        })
    })
    describe("when two input arrays are not empty", () => {
        it("returns an array of matching elements", () => {
            let arr1 = ["Toyota", "Honda", "Nissan", "BMW", "Chevy", "Ford"];
            let arr2 = ["Ford", "BMW", "Audi", "Mercedes"];
            let result = searchMatch(arr1, arr2);
            expect(result).toEqual(["BMW", "Ford"]);
        });

        describe("when no match is found", () => {
            it("returns an empty array", () => {
                let arr1 = ["Toyota", "Honda", "Nissan", "BMW", "Chevy", "Ford"];
                let arr2 = ["Audi", "Mercedes", "Tesla"];
                let result = searchMatch(arr1, arr2);
                expect(result).toEqual([]);
            })
        })
    })
})
