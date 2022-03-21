const { expect, describe, it } = require('@jest/globals');
const sum = require('./main');

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

    it('adds numbers', () => {
        expect(sum(1, 2)).toBe(3);
    })
})