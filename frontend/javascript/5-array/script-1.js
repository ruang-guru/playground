// Coba tuliskan 5 operasi array berikut
//
// 1. Buatlah array dengan elemen-elemen sebagai berikut: "Ant", "Bee"
// 2. Tambahkan elemen "Cow" ke dalam array tersebut.
// 3. Ubah elemen yang ada di tengah dengan "Owl". Kode yang dibuat untuk mencari nilai yang ada di tengah harus bekerja untuk semua array dengan jumlah ganjil. Misal jumlah array 7, maka "Owl akan ditempatkan pada indeks ke 4".
// 4. Hapus elemen pertama pada array.
// 5. Tambahkan "Eagle" dan "Elephant" di bagian depan dari array.

let animals = ["Ant", "Bee"];
animals.push("Cow");
animals[Math.floor((animals.length - 1) / 2)] = "Owl";
animals.shift();
animals.unshift("Eagle", "Elephant");

console.log(animals);