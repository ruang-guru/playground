// lakukan call pada API ini https://jsonplaceholder.typicode.com/todos/1
// dan console.log hasilnya menggunakan async await
// console log error apabila terjadi error

// karena ini adalah environtment node-js, fetch tidak bisa digunakan
// tuliskan program kalian dibawah, dan test kode tersebut di
// https://playcode.io/new/

const callApi = async () => {
  // kerjakan disini
  //beginanswer
  try {
    const res = await (
      await fetch("https://jsonplaceholder.typicode.com/todos/1")
    ).json();

    console.log(res);
  } catch (error) {
    console.log(error);
  } finally {
    console.log("API selesai dipanggil");
  }
  //endanswer
};

callApi();
