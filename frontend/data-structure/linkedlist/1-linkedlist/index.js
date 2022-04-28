// Tulis sebuah fungsi untuk menghapus sebuah node dalam sebuah single-linked list.
// Kita tidak akan diberikan akses ke head dari list, melainkan kita akan diberikan akses ke node yang akan dihapus secara langsung.
// Dapat dipastikan node yang akan dihapus bukan merupakan tail node dalam list.
//
// Contoh 1
// Input: head = [4,5,1,9], node = 5
// Output: [4,1,9]
// Explanation: Kita mengetahui node 5 akan dihapus, sehingga linked list menjadi 4 -> 1 -> 9 setelah memanggil fungsi delete.
//
// Contoh 2
// Input: head = [4,5,1,9], node = 1
// Output: [4,5,9]
// Explanation: Kita mengetahui node 5 akan dihapus, sehingga linked list menjadi 4 -> 5 -> 9 setelah memanggil fungsi delete.

function ListNode(val) {
  this.val = val;
  this.next = null;
}

let deleteNode = function(node) {
  node.val = node.next.val;
  node.next = node.next.next;
};
