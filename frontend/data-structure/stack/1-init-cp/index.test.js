const { expect, describe, it } = require('@jest/globals');
const Stack = require('./index');

describe("Stack", () => {
    describe("when init new stack", () => {
        it("should return stack with empty array, size equal to 10, and top equal to -1", () => {
            let stack = new Stack();
            expect(stack.data.length).toEqual(0);
            expect(stack.size).toEqual(10);
            expect(stack.top).toEqual(-1);
        });
    })
})
