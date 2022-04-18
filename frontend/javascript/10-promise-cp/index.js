/**
 * Fungsi requestStarWarsPeopleById, menggunakan teknik callback untuk mengembalikan datanya.
 * Ubahlah fungsi tersebut menjadi bentuk Promise.
 *
 * Catatan:
 * Sedikit penjelasan pada node.js untuk melakukan http request menggunakan module request,
 * oleh karena itu di sini kita menggunakan module https.
 *
 */
const https = require("https");

function requestStarWarsPeopleById(peopleId, onReturn, onError) {
  https
    .get("https://swapi.dev/api/people/" + peopleId, (res) => {
      let result = "";

      if (res.statusCode !== 200) {
        onError(new Error(res.statusCode));
      }

      res.on("data", (d) => {
        result += d;
      });

      res.on("end", () => {
        onReturn(JSON.parse(result));
      });
    })
    .on("error", (e) => {
      onError(e);
    });
}

function promiseStarWarsPeopleById(peopleId) {
  // TODO: answer here
}

module.exports = {
  promiseStarWarsPeopleById,
  requestStarWarsPeopleById,
};
