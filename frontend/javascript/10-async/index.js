/**
 * Fungsi getStarWarsData merupakan sebuah fungsi yang mengembalikan promise.
 * Yang mana promise tersebut akan mengembalikan data dari swapi tergantung parameter request yang diberikan.
 *
 * Sebelumnya kita sudah belajar menggunakan cara penggunaan promise. 
 * 
 * Kali ini kita akan melihat bagaimana perbedaan ketika menggunakan async.
 */

const https = require("https");

function getStarWarsData(url) {
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
          resolve(result);
        });
      })
      .on("error", (e) => {
        reject(e);
      });
  });
}

function getPromisePeopleById(id) {
  return getStarWarsData("https://swapi.dev/api/people/" + id).then(
    (stringResult) => {
      return JSON.parse(stringResult);
    }
  );
}

async function getDataPeopleById(id) {
  const stringResult = await getStarWarsData(
    "https://swapi.dev/api/people/" + id
  );

  return JSON.parse(stringResult);
}

function accessPeoplePromise() {
  getPromisePeopleById(1).then((people) => console.log(people));
}
accessPeoplePromise();

async function accessPeople() {
  const people = await getDataPeopleById(1);
  console.log(people);
}
accessPeople();
