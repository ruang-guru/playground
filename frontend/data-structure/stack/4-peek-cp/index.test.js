const { expect, describe, it } = require('@jest/globals');
const Stack = require('./index');

describe("Stack", () => {
    describe("when peek on empty stack", () => {
        it("should throw stack is empty", () => {
            let stack = new Stack();

            expect(() => {
                stack.peek();
            }).toThrow('stack is empty');
            expect(stack.data.length).toEqual(0);
            expect(stack.top).toEqual(-1);
        });
    });
    describe("when peek on non-empty stack", () => {
        it("should return most top element", () => {
            let stack = new Stack();
            stack.push(1);
            stack.push(2);
            stack.push(3);
            stack.push(4);
            stack.push(5);

            let el = stack.peek()
            expect(el).toEqual(5)
            expect(stack.data.length).toEqual(5);
            expect(stack.top).toEqual(4);
        });
    });
})
