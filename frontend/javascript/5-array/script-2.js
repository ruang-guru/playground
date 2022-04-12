// Menghapus elemen dalam array yang duplikat
// contoh: 
// ["ant", "bee", "bee", "cow", "ant", "cow", "elephant", "dog", "elephant"]
// hasil:
// ["ant", "bee", "cow", "elephant", "dog"]

let animals = ["ant", "bee", "bee", "cow", "ant", "cow", "elephant", "dog", "elephant"];
  
let unique_animals = [];
animals.forEach((animal) => {
    if (!unique_animals.includes(animal)) {
        unique_animals.push(animal);
    }
  });
  
console.log(unique_animals);