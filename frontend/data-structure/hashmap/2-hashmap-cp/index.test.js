const { expect, describe, it } = require('@jest/globals');
const { anagramChecker } = require('./index');

describe("anagramChecker", () => {
    describe("when the two input strings have different length", () => {
        it("should return false", () => {
            let str1 = "makan";
            let str2 = "namaka";
            expect(anagramChecker(str1, str2)).toBe(false);
        })
    })
    describe("when the two input strings have the same length", () => {
        describe("when the two input strings have the same characters", () => {
            it("should return true", () => {
                let str1 = "fried";
                let str2 = "fired";
                expect(anagramChecker(str1, str2)).toBe(true);

                str1 = "keen";
                str2 = "knee";
                expect(anagramChecker(str1, str2)).toBe(true);
            })
        })
        describe("when the two input strings have different characters", () => {
            it("should return false", () => {
                let str1 = "apple";
                let str2 = "aplle";
                expect(anagramChecker(str1, str2)).toBe(false);
            })
        })
    })
})
