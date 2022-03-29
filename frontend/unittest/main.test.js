const { expect, describe, it } = require('@jest/globals');
const { sum, multiply} = require('./main');

describe('adds', () => {
    describe("when both numbers are positive", () => {
        it('returns a positive number', () => {
            expect(sum(1, 2)).toBeGreaterThan(0); 
        })
    })
    describe("when both numbers are negative", () => {
        it('returns a negative number', () => {
            expect(sum(-1, -2)).toBeLessThan(0); 
        })
    })

    describe("when both numbers are zero", () => {
        it('returns zero', () => {
            expect(sum(0, 0)).toBe(0); 
        })
    })

    describe("when one of the number is 0", () => {
        it("returns the other number", () => {
            expect(sum(0, 1)).toBe(1);
            expect(sum(1, 0)).toBe(1);
        })
    })

    it('adds numbers', () => {
        expect(sum(1, 2)).toBe(3);
    })
})

describe("multiply", () => {
    describe("when one of the number is 0", () => {
        it('returns 0', () => {
            expect(multiply(1, 0)).toBe(0);
            expect(multiply(0, 1)).toBe(0);
        })
    })
    describe("when one of the number is 1", () => {
        it("returns the other number", () => {
            expect(multiply(1, 2)).toBe(2);
            expect(multiply(2, 1)).toBe(2);
        })
    })

    it('multiply the numbers', () => {
        expect(multiply(2, 3)).toBe(6);
    })
})