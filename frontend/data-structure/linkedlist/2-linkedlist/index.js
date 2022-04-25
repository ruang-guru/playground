// Diberikan head dari singly linked list, balikkan (reverse) list, dan kembalikan daftar yang sudah dibalik.
//
// Contoh 1
// Input: head = [1,2,3,4,5]
// Output: [5,4,3,2,1]
//
// Contoh 2
// Input: head = [1,2]
// Output: [2,1]
//
// Contoh 3
// Input: head = []
// Output: []

function ListNode(val, next) {
    this.val = (val===undefined ? 0 : val)
    this.next = (next===undefined ? null : next)
}

let reverseList = function(head) {
    let prev = null, next = null;
    while (head) {
        next = head.next
        head.next = prev;
        prev = head;
        head = next;
    }
    return prev;
};