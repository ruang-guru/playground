# DEMO 1 - Buat project React tanpa create-react-app

Demo ini adalah demo untuk membuat project react dari awal tanpa menggunakan create-react-app

## Setting up React environment

- Masuk ke folder `frontend\react\react-set-up\manual-set-up` masing-masing
- Untuk inisalisasi project baru lakukan `npm init` atau `yarn init` untuk inisialisasi project baru

```shell
contoh D:\LEARNING\RUANGGURU\frontend\react\react-set-up\manual-set-up> npm init
```

```json
{
  "name": "manual-set-up",
  "version": "1.0.0",
  "description": "manually set up react app without RCA",
  "main": "index.js",
  "scripts": {
    "test": "echo \"Error: no test specified\" && exit 1"
  },
  "author": "",
  "license": "ISC"
}
```

- Lakukan langkah-langkah berikut dalam membuat project baru react

1. Install webpack dependencies

```shell
contoh D:\LEARNING\RUANGGURU\frontend\react\react-set-up\manual-set-up> npm i webpack webpack-cli webpack-dev-server
```

- **webpack** (<https://webpack.js.org/>) merupakan sebuah module bundler JavaScript untuk membungkus module/dependency/library/assets menjadi **satu paket file** berekstensi JavaScript (.js).
- **webpack-cli** merupakan package yang digunakan untuk membantu kita menjalankan webpack melalui sebuah perintah (Command Line Interface).
- **webpack-dev-server** memudahkan kita dalam proses pengembangan karena kita dapat menjalankan fitur live-reloading

2. Install Babel Dependecies

```shell
contoh D:\LEARNING\RUANGGURU\frontend\react\react-set-up\manual-set-up> npm i babel-loader @babel/preset-env @babel/core @babel/preset-react @babel/cli
```

- **Babel** merupakan sebuah transpiler yang bertugas untuk mengubah sintaks JavaScript modern (ES6+) menjadi sintaks yang dapat didukung penuh oleh seluruh browser
- **@babel/core** merupakan package inti yang harus dipasang ketika kita hendak menggunakan babel
- **babel-loader** merupakan package yang diperlukan untuk menggunakan babel sebagai loader pada webpack
- **@babel/preset-env** merupakan package preset yang akan kita gunakan untuk membantu babel-loader dalam melakukan tugasnya
- **@babel/preset-react** membantu dalam mengonversi file html ke file berbasis react
- **@babel/cli** command line interface untuk menggunakan babel

3. Install react and react-dom

```shell
npm i react react-dom
```

4. Buatlah direktori **public** yang akan menyimpan aset-aset statis dan file **index.html** yang akan digunakan untuk merender aplikasi kita

```shell
contoh D:\LEARNING\RUANGGURU\playground\frontend\react\react-set-up\manual-set-up> mkdir public
contoh D:\LEARNING\RUANGGURU\playground\frontend\react\react-set-up\manual-set-up> cd public
contoh D:\LEARNING\RUANGGURU\playground\frontend\react\react-set-up\manual-set-up\public> touch index.html
```

- pada file index.html tambahkan code berikut

```html
<!DOCTYPE html>
<html lang="en">
  <head>
    <meta charset="UTF-8" />
    <title>React App</title>
  </head>
  <body>
    <div id="root"></div>
    <script src="main.js"></script>
  </body>
</html>
```

5. Buatlah direktori **src** pada folder **react-set-up** untuk menyimpan _page, component_ atau sumber daya lainnya

```shell
contoh D:\LEARNING\RUANGGURU\playground\frontend\react\react-set-up\manual-set-up> mkdir src
contoh D:\LEARNING\RUANGGURU\playground\frontend\react\react-set-up\manual-set-up> cd src
```

6. Buatlah file **App.js** pada direktori **src**. **App.js** adalah file untuk **App Component**. **App Component** adalah komponen utama dalam React yang bertindak sebagai wadah untuk semua komponen lainnya.

```shell
contoh D:\LEARNING\RUANGGURU\playground\frontend\react\react-set-up\manual-set-up\src> touch App.js
```

- Masukkan code berikut ke dalam file **App.js**

```js
import React from "react";

const App = () => {
  return <h1>Hello React</h1>;
};

export default App;
```

7. Masih pada direktori **src** buatlah file **index.js** yang akan menghubungkan file **App.js** kita dengan **index.html**

```shell
contoh D:\LEARNING\RUANGGURU\playground\frontend\react\react-set-up\manual-set-up\src> touch index.js
```

- Lalu masukkan kode berikut

```js
import React, { StrictMode } from "react";
import { createRoot } from "react-dom/client";
import App from "./App";

const rootElement = document.getElementById("root");
const root = createRoot(rootElement);

root.render(
  <StrictMode>
    <App />
  </StrictMode>
);
```

8. Buatlah file **webpack.config.js** di dalam `frontend\react\react-set-up\manual-set-up` dan install package `path`

```shell
contoh D:\LEARNING\RUANGGURU\frontend\react\react-set-up\manual-set-up> touch webpack.config.js
contoh D:\LEARNING\RUANGGURU\frontend\react\react-set-up\manual-set-up> npm i path
```

- **_Paste code_** berikut ke file **webpack.config.js**

```js
const path = require("path");

/*We are basically telling webpack to take index.js from entry. Then check for all file extensions in resolve. 
After that apply all the rules in module.rules and produce the output and place it in main.js in the public folder.*/

module.exports = {
  /** "mode"
   * the environment - development, production, none. tells webpack
   * to use its built-in optimizations accordingly. default is production
   */
  mode: "development",
  /** "entry"
   * the entry point
   */
  entry: ".=/src/index.js",
  output: {
    /** "path"
     * the folder path of the output file
     */
    path: path.resolve(__dirname, "/public"),
    /** "filename"
     * the name of the output file
     */
    filename: "main.js"
  },
  /** "target"
   * setting "node" as target app (server side), and setting it as "web" is
   * for browser (client side). Default is "web"
   */
  target: "web",
  devServer: {
    /** "port"
     * port of dev server
     */
    port: "9500",
    /** "static"
     * This property tells Webpack what static file it should serve
     */
    static: ["./public"],
    /** "open"
     * opens the browser after server is successfully started
     */
    open: true,
    /** "hot"
     * enabling and disabling HMR. takes "true", "false" and "only".
     * "only" is used if enable Hot Module Replacement without page
     * refresh as a fallback in case of build failures
     */
    hot: true,
    /** "liveReload"
     * disable live reload on the browser. "hot" must be set to false for this to work
     */
    liveReload: true
  },
  resolve: {
    /** "extensions"
     * If multiple files share the same name but have different extensions, webpack will
     * resolve the one with the extension listed first in the array and skip the rest.
     * This is what enables users to leave off the extension when importing
     */
    extensions: [".js", ".jsx", ".json"]
  },
  module: {
    /** "rules"
     * This says - "Hey webpack compiler, when you come across a path that resolves to a '.js or .jsx'
     * file inside of a require()/import statement, use the babel-loader to transform it before you
     * add it to the bundle. And in this process, kindly make sure to exclude node_modules folder from
     * being searched"
     */
    rules: [
      {
        test: /\.(js|jsx)$/, //kind of file extension this rule should look for and apply in test
        exclude: /node_modules/, //folder to be excluded
        use: "babel-loader" //loader which we are going to use
      }
    ]
  }
};
```

- Selanjutnya, buat file **.babelrc** dan masukkan code dibawah ini

```js
{
  /*
    a preset is a set of plugins used to support particular language features.
    The two presets Babel uses by default: es2015, react
  */
  "presets": [
    "@babel/preset-env", //compiling ES2015+ syntax
    "@babel/preset-react" //for react
  ]
}

```

- Update **package.json** file. Tambahkan script `start`

```json
{
  "name": "manual-set-up",
  "version": "1.0.0",
  "description": "manually set up react app without RCA",
  "main": "index.js",
  "scripts": {
    "test": "echo \"Error: no test specified\" && exit 1",
    "start": "webpack-dev-server ./src/index.js"


  },

```

- Tambahkan script `"start": "webpack-dev-server ./src/index.js"`

- Jalankan script `npm start` dari terminal untuk menjalankan react dev server kita

```shell
contoh D:\LEARNING\RUANGGURU\frontend\react\react-set-up\manual-set-up> npm start
```
