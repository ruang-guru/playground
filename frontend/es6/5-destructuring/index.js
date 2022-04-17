/**
 * Fungsi yang hanya menerima array dengan dua element, kemudian mengalikan kedua elemen tersebut
 */
const multiplyArray = (arr) => {
  // Pastikan hanya menerima dua element.
  if (arr.length > 2) {
    throw new Error("too many element");
  }
  // Akses kedua elemen menggunakan destructuring assignment.
  const [x, y] = arr;

  // Kalikan kedua element tersebut.
  return x * y;
};

/**
 * Fungsi yang menerima object `people`, dan mengembalikan string ucapan perkenalan
 * `{name, job}` menggunakan destructuring assignment untuk mengakses properti dari objek `people`.
 */
const greetings = ({ name, job }) => `Hello my name is ${name}, I am a ${job}`;
