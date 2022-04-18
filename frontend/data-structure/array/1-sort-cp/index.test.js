const { expect, describe, it } = require('@jest/globals');
const { sort, rotateLeft } = require('./index');

describe('sort', () => {
    describe('when the array is already sorted', () => {
        it("returns the same array", () => {
            let nums = [1, 2, 3, 4, 5];
            let sorted = sort(nums);
            expect(sorted).toEqual(nums);
        })
    })
    describe("when the array is not sorted", () => {
        it("returns the sorted array", () => {
            let nums = [4, 5, 2, 1, 3];
            let sorted = sort(nums);
            expect(sorted).toEqual([1, 2, 3, 4, 5]);
        })
    })
})

describe("rotateLeft", () => {
    describe("when the array is empty", () => {
        it("returns the same array", () => {
            let nums = [];
            let rotated = rotateLeft(4, nums);
            expect(rotated).toEqual(nums);
        })
    })

    describe("when the array is not empty", () => {
        it("returns the rotated array", () => {
            let nums = [1,2,3,4,5];
            let rotated = rotateLeft(4, nums);
            expect(rotated).toEqual([5,1,2,3,4]);
        })
    });
})