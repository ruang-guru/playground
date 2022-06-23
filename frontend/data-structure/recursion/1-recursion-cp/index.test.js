const { expect, describe, it } = require('@jest/globals');
const { myPow } = require('./index');

describe("myPow", () => {
    describe("when the power is 0", () => {
        it("returns 1", () => {
            let x = 20;
            let result = myPow(x, 0);
            expect(result).toEqual(1);
        })
    })

    describe("when the power > 0", () => {
        it("returns the correct value", () => {
            let x = 3;
            let n = 2;
            let result = myPow(x, n);
            expect(result).toEqual(9);
        })
    });

    describe("when the power < 0", () => {
        it("returns the correct value", () => {
            let x = 2;
            let n = -3;
            let result = myPow(x, n);
            expect(result).toEqual(0.125);
        })
    });
})