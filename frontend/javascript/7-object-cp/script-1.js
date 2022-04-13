// Dari kode di bawah ini

let programming = {
    languages: ["JavaScript", "Python", "Ruby"],
    isChallenging: true,
    isRewarding: true,
    difficulty: 8,
    jokes: "http://stackoverflow.com/questions/234075/what-is-your-best-programmer-joke"
};

// buat kode untuk manipulasi data dalam objek berdasarkan instruksi berikut

// 1. Tulis perintah untuk menambahkan bahasa "Go" ke akhir array languages.
// 2. Ubah difficulty menjadi 7.
// 3. Menggunakan kata kunci delete, tulis perintah untuk menghapus key jokes dari objek programming.
// 4. Tulis perintah untuk menambahkan key baru yang disebut isFun dan nilai true ke objek programming.
// 5. Menggunakan loop, iterasi array languages dan console.log semua bahasa.
// 6. Menggunakan loop, console.log semua key dalam objek pemrograman.
// 7. Menggunakan loop, console.log semua value dalam objek pemrograman.

// beginanswer
// 1
programming.languages.push("Go");

// 2
programming.difficulty = 7;

// 3
delete programming.jokes;

// 4
programming.isFun = true;

// 5
for (let i = 0; i < programming.languages.length; i++) {
  console.log(programming.languages[i]);
}

// or 
for (let language of programming.languages) {
  console.log(language);
}

// 6
for (let key in programming){
  console.log(key);
}

// 7
for (let key in programming){
  console.log(programming[key]);
}
// endanswer