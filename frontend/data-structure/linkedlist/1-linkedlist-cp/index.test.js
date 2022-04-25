const { expect, describe, it } = require('@jest/globals');
const { ListNode, middleNode } = require('./index');

describe('middleNode', () => {
    describe('when the list is odd', () => {
        it("returns the correct middle node", () => {
            let list = ListNode(1, ListNode(2, ListNode(3, null)));
            let middle = middleNode(list);
            expect(middle).toEqual(ListNode(2, ListNode(3, null)));
        })
    })
})