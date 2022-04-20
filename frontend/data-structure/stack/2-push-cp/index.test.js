const { expect, describe, it } = require('@jest/globals');
const Stack = require('./index');

describe("Stack", () => {
    describe("when push element on empty stack", () => {
        it("should append stack data with element", () => {
            let stack = new Stack();
            stack.push(1);
            expect(stack.data.length).toEqual(1);
            expect(stack.top).toEqual(0);
            expect(stack.data[0]).toEqual(1);
        });
    });
    describe("when push element on full stack", () => {
        it("should return stack overflow", () => {
            let stack = new Stack();
            stack.push(1);
            stack.push(2);
            stack.push(3);
            stack.push(4);
            stack.push(5);
            stack.push(6);
            stack.push(7);
            stack.push(8);
            stack.push(9);
            stack.push(10);

            expect(() => {
                stack.push(11);
            }).toThrow('stack overflow');
            expect(stack.data.length).toEqual(10);
            expect(stack.top).toEqual(9);
        });
    });
})
