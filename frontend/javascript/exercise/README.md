# Javascript Exercise

Sebelum mengerjakan soal dibawah ini pastikan kalian sudah melakukan instalasi **Yarn** dan **Jest** sebelumnya.
Jika kalian belum melakukan instalasi, silahkan baca terlebih dahulu instruksi dibawah.

## Demo:

- [Zoo Ticket](zoo-ticket/main.js)
- [Sentence](sentence/main.js)
- [Index Accessing](index-accessing/main.js)
- [Braking Sentence](breaking-sentence/main.js)
- [Graduate Student](graduate-student/main.js)
- [University](university/main.js)
- [I love coding](i-love-coding/main.js)
- [Besar Kecil](besar-kecil/main.js)
- [Jumlah Kata](jumlah-kata/main.js)
- [Palindrome](palindrome/main.js)
- [XOXO](xoxo/main.js)

## Practice:

- [Braking Sentence](breaking-sentence-cp/main.js)
- [Gatcha](gatcha-cp/main.js)
- [Cafe Visitation](cafe-visitation-cp/main.js)
- [Sandi](sandi-cp/main.js)
- [Fix Data](fix-data-cp/main.js)
- [Konversi Menit](konversi-menit-cp/main.js)
- [Palindrome Angka](palindrome-angka-cp/main.js)
- [Ganjil Genap](ganjil-genap-cp/main.js)

---

## How to install Yarn

Disarankan untuk menginstal Yarn melalui manajer paket npm, yang dibundel dengan Node.js saat ingin melakukan instalasi di dalam sistem.

Setelah melakukan instalasi npm, kita dapat menjalankan perintah berikut untuk menginstal Yarn:

```shell
npm install --global yarn
```

Atau kita juga bisa menjalankan perintah berdasarkan OS yang kita gunakan.

### Ubuntu

Di Debian atau Ubuntu Linux, Kita dapat menginstal Yarn melalui repositori Debian package.
Kita harus terlebih dahulu mengonfigurasi repositori:

```shell
curl -sS https://dl.yarnpkg.com/debian/pubkey.gpg | sudo apt-key add -
echo "deb https://dl.yarnpkg.com/debian/ stable main" | sudo tee /etc/apt/sources.list.d/yarn.list
```

Jika sudah, pada terminal kita bisa menjalankan:

```shell
sudo apt update && sudo apt install yarn
```

### macOS

Kita dapat menginstal Yarn melalui manajer paket Homebrew.
Ini juga akan sekaligus menginstal Node.js jika belum diinstal.

```shell
brew install yarn
```

Atau melalui MacPort dengan cara berikut:

```shell
sudo port install yarn
```

### Windows

Pada OS windows, kalian dapat melakukan instalasi dengan mendownlad file `.msi` berikut:
[Download disini](https://classic.yarnpkg.com/latest.msi)

Pastikan sebelum menginstal file diatas, kalian harus menginstal Node.js terlebih dahulu

Atau melalui Chocolatey dengan cara berikut:

```shell
choco install yarn
```

## Check installation

```shell
yarn --version
```

---

## How to install Jest

Untuk menginstal Jest, kalian dapat menggunakan Yarn atau npm dengan cara:

- Install Jest menggunakan yarn

  ```shell
  yarn add jest
  ```

- Install Jest menggunakan npm
  ```shell
  npm install jest
  ```
