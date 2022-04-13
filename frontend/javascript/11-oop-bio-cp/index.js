/**
 * Pada soal kali ini disediakan class `Person` dimana kalian diminta untuk melengkapi serta membuat beberapa hal dibawah ini:
 * - Class Person yang memiliki properti name, age, dan job
 * - Buatlah getter untuk masing masing Properti diatas dengan format `getName` untuk mendapatkan nama, `getAge` untuk mendapatkan usia, dan `getJob` untuk mendapatkan pekerjaan dari class Person.
 * - Buatlah setter untuk masing masing Properti diatas dengan format `setName` untuk mengubah nilai nama, `setAge` untuk mengubah nilai usia, dan `setJob` untuk mengubah nilai pekerjaan dari class Person
 */

class Person {
  constructor(name, age, job) {
    this.name = name;
    this.age = age;
    this.job = job;
  }
  
  //beginanswer
  get getName() {
    return this.name;
  }

  get getAge() {
    return this.age;
  }

  get getJob() {
    return this.job;
  }

  set setName(name) {
    this.name = name;
  }

  set setAge(age) {
    this.age = age;
  }

  set setJob(job) {
    this.job = job;
  }
  //endanswer
}

module.exports = Person