// Teman kalian berjanji akan membuatkan kalian kue ulang tahun.
// Kalau semuanya berjalan lancar dan dia tidak sakit, maka kalian akan mendapat kue.

const onMyBirthday = (isMyFriendSick) => {
  return new Promise((resolve, reject) => {
    setTimeout(() => {
      if (!isMyFriendSick) {
        return resolve(true);
      }

      return reject(false);
    }, 2000);
  });
};

onMyBirthday(false)
  .then((_) => console.log("Dapat kue"))
  .catch((_) => console.log("Tidak dapat kue"));

// underscore pada then dan catch menandakan bahwa hasil dari promise itu sendiri tidak dipakai
// contoh dibawah untuk yang dipakai

const onMyNextBirthday = (isMyFriendSick, numberOfCakes) => {
  return new Promise((resolve, reject) => {
    setTimeout(() => {
      if (!isMyFriendSick) {
        return resolve(numberOfCakes);
      }

      return reject(0);
    }, 2000);
  });
};

onMyNextBirthday(true, 3)
  .then((res) =>
    console.log(`Teman saya tidak sakit, dia membuatkan saya ${res} kue`)
  )
  .catch((err) =>
    console.log(`Teman saya sakit, sehingga saya mendapat ${err} kue`)
  );
