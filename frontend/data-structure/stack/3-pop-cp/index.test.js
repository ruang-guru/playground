const { expect, describe, it } = require('@jest/globals');
const Stack = require('./index');

describe("Stack", () => {
    describe("when pop on empty stack", () => {
        it("should throw stack underflow", () => {
            let stack = new Stack();

            expect(() => {
                stack.pop();
            }).toThrow('stack underflow');
            expect(stack.data.length).toEqual(0);
            expect(stack.top).toEqual(-1);
        });
    });
    describe("when pop on non-empty stack", () => {
        it("should return most top element", () => {
            let stack = new Stack();
            stack.push(1);
            stack.push(2);
            stack.push(3);
            stack.push(4);
            stack.push(5);

            let el = stack.pop()
            expect(el).toEqual(5)
            expect(stack.data.length).toEqual(4);
            expect(stack.top).toEqual(3);
        });
    });
})
