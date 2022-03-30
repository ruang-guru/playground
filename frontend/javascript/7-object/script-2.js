// Dari data di bawah, buatlah object nya, dan buatlah kode untuk menampilkan apakah sudah membaca/belum berdasarkan readingStatus nya
//
// Data:
// 1. author: 'Bill Gates', title: 'The Road Ahead', readingStatus: true
// 2. author: 'Steve Jobs', title: 'Walter Isaacson', readingStatus: true
// 3. author: 'Suzanne Collins', title:  'Mockingjay: The Final Book of The Hunger Games', readingStatus: false
//
// contoh output:
// 'The Road Ahead' by Bill Gates sudah dibaca.
// 'Walter Isaacson' by Steve Jobs sudah dibaca.
// 'Mockingjay: The Final Book of The Hunger Games' by Suzanne Collins belum dibaca.

var library = [ 
    {
        author: 'Bill Gates',
        title: 'The Road Ahead',
        readingStatus: true
    },
    {
        author: 'Steve Jobs',
        title: 'Walter Isaacson',
        readingStatus: true
    },
    {
        author: 'Mockingjay: The Final Book of The Hunger Games',
        title: 'Suzanne Collins',
        readingStatus: false
    }
];

for (var i = 0; i < library.length; i++) {
  var book = "'" + library[i].title + "'" + ' by ' + library[i].author;
  if (library[i].readingStatus) {
    console.log(book + "sudah dibaca.");
  } else {
    console.log(book + "belum dibaca.");
  }
}