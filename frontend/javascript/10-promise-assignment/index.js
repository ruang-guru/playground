/**
 * Buatlah fungsi untuk mengambil data karakter pada film starwars, beserta semua film yang dimainkan oleh karakter tersebut.
 *
 * Contoh: people/1/
 * Output yang harapkan:
 * {
 *  "name": "Luke Skywalker",
 *  "height": "172",
 *  "mass": "77",
 *  "hair_color": "blond",
 *  "skin_color": "fair",
 *  "eye_color": "blue",
 *  "birth_year": "19BBY",
 *  "gender": "male",
 *  "films": [
 *    {
 *      "title": "The Empire Strikes Back",
 *      "episode_id": 5,
 *    },
 *    {
 *      "title": "Revenge of the Sith",
 *      "episode_id": 3,
 *     }
 *  ]
 *
 * Catatan:
 * - Lihatlah dan coba pahami dokumentasi dari endpont berikut: https://swapi.dev/documentation#people dan https://swapi.dev/documentation#films
 * - Kedua endpoint tersebut akan meliki semua data yang dibutuhan dari output yang diharapkan dari fungsi ini.
 * - Pastikan fungsi yang kalian buat mengembalikan format data yang sesuai dengan yang diharapkan, jika tidak test akan error.
 * - Untuk mempermudah perulangan kalian bisa menggunakan perintah Promise.all(), https://developer.mozilla.org/en-US/docs/Web/JavaScript/Reference/Global_Objects/Promise/all
 */

const https = require("https"); // Dibutuhkan sebagai protokol untuk akses data.

/**
 * Untuk memudahkan kami telah menyiapkan fungsi untuk melakukan request data starwars.
 * Contoh params: https://swapi.dev/api/people/1
 */

function promiseStarWarsData(url) {
  return new Promise((resolve, reject) => {
    https
      .get(url, (res) => {
        let result = "";

        if (res.statusCode !== 200) {
          reject(new Error(res.statusCode));
        }

        res.on("data", (d) => {
          result += d;
        });

        res.on("end", () => {
          resolve(JSON.parse(result));
        });
      })
      .on("error", (e) => {
        reject(e);
      });
  });
}

function getDataPeopleByIdWithFilms(peopleId) {
  // TODO: answer here
}

module.exports = { getDataPeopleByIdWithFilms };
