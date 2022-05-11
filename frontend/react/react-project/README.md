# React Project "Instagram Clone"

## Setup
Jalankan `npm install` untuk meng-install dependency yang dibutuhkan

## BackEnd API
Dokumentasi API dapat dilihat di [https://hackmd.io/@RieGan/rg-km_insta-clone_api](https://hackmd.io/@RieGan/rg-km_insta-clone_api)

## Test
Test file berada pada directory `src/__test__/...`, untuk menjalankan test tersebut dapat dilakukan dengan command

```shell
npm test
```
atau
```shell
npm test -- <nama_file_test>
```
untuk menjalankan satu persatu.

## Assignment
### First Render
__file test :__ `1-first-render.test.js` \
__file component :__ `src/App.js`

Spesifikasi:
- Container dengan atribut `aria-label="App"`
- Title atau text dengan atribut `aria-label="App Title"`

### Like & Dislike Component
__file test :__ `3-like-button-component.test.js` \
__file component :__ `src/components/LikeDislikeButton.js`

Spesifikasi:
- Component memiliki props 
    ```JSON
    {
        id: , // ID dari post
        image: , // URL post image
        isLiked: , // flag apakah sudah di-like oleh user
        isDisliked: , // flag apakah sudah di-dislike oleh user
        likeCount: , // jumlah like
        dislikeCount:  // jumlah dislike
    }
    ```
- Button Like dengan atribut `aria-label="Like Button"`
- Button Dislike dengan atribut `aria-label="Dislike Button"`
- Text untuk menampilkan jumlah like dengan atribut `aria-label="Like Count"`
- Text untuk menampilkan jumlah dislike dengan atribut `aria-label="Dislike Count"`
- apabila user melakukan click like, jumlah like akan bertambah
- apabila user melakikan click dislike, jumlah dislike akan bertambah
- apabila user melakukan click like setelah like, maka button akan melakukan unlike begitu pula dengan dislike
- apabila user melakukan click like setelah dislike, maka status dislike akan berubah menjadi like, begitu pula sebaliknya

### Navbar Component
__file test :__ `3-navbar-component.test.js` \
__file component :__ `src/components/Navbar.js` 

Spesifikasi: 
- Container component dengan atribut `aria-label="Navbar"`
- Terdapat text link menuju "/" dengan atribut `ria-label="App Title"`
- Terdapat gambar logo dengan atribut `aria-label="App Logo"`

### Post Card Component
__file test :__ `3-post-card-component.test.js` \
__file component: __`src/components/PostCard.js` 

Spesifikasi:
- Container component dengan atribut `aria-label="Post Card"`
- Component memiliki props
  ```JSON
  {
      image: , // URL gambar
      caption: , // caption dari post
      username: , // nama pemilik post
      userId: , // user id dari pemilik post
      date: , // tanggal upload
  }
  ```
- terdapat gambar dengan atribut `aria-label="Post Image"`
- terdapat nama pemiliki post dengan atribut `aria-label="Post User Name"`
- terdapat caption dengan atribut `aria-label="Post Caption"`
- terdapat tanggal dengan atribut `aria-label="Post Date"`

### Login
__file test :__ `4-api-auth.test.js` 

Spesifikasi: 
- Component navbar terdapat button login dengan atribut `aria-label="Login"` (muncul ketika user belum login)
- Gunakan function `auth` yang terdapat pada `src/api/auth.js` untuk melakukan login
- Gunakan function `getSession` yang terdapat pada `src/api/auth.js` untuk mendapatkan informasi user yang ter-loggedin
- apabila user telah login (memiliki session) tampilkan nama profile dengan atribut `aria-label="Profile"`

### Like Dislike API
__file test :__ `4-api-like-dislike.test.js` 

Spesifikasi: 
- Component like dislike button dapat memanggil API

### Form Upload
__file test :__ `4-form-upload.test.js` \
__file component :__ `src/components/UploadForm.js` 

Spesifikasi:
- Container component memiliki atribut `aria-label="Upload Form"`
- Terdapat input untuk caption dengan atribut `aria-label="Caption Input"`
- Terdapat input file gambar dengan atribut `aria-label="Image Input"`
- file gambar yang bisa diterima hanya `.jpg`, `.png`, dan `.gif`
- Terdapat button untuk submit dengan atribut `aria-label="Submit Button"`
- Submit button akan melakukan request POST ke API dengan menggunakan encoding `multipart/form-data` dengan isian 
  ```JSON
  {
      content: , //caption
      image: , //file image
  }
  ```

### Posts API with Components Integration
__file test :__ `4-api-childrens-component.test.js` 

Spesifikasi:
- Component App (component utama) terdapat component `Navbar`, `PostCard`, dan `UploadForm`
- Gunakan API untuk mendapatkan list post yang terupload oleh semua user
- Component `PostCard` digunakan untuk menampilkan post post yang didapatkan dari API
- Apabila user berhasil melakukan upload, otomatis akan tertambahkan di halaman tanpa melakukan reload

### Session Context
__file test :__ `6-user-context.test.js` \
__file component :__ `src/context/SessionContext.js` 

Spesifikasi: 
- Context memiliki value 
  ```JSON
  {
    user: {
        id, // user id
        name, // user name
        image, // user image
    },
    setUser, // function untuk mengubah user
    isLoggedIn, // status apakah user telah login
    setIsLoggedIn, // function untuk mengubah status isLoggedIn
  }
  ```
- Provider yang memberikan semua value tersebut

### Router
__file test :__ `8-react-router.test.js` \
__file component:__ `src/App.js`, `src/components/Profile.js` 
Spesifikasi:
- gunakan BrowserRouter
- terdapat index page "/" (halaman yang menampilkan semua post + form upload)
- terdapat profile page "/profile" (halaman yang menampilkan semua post dari user berdasarkan userID tersebut)
- page profile mendapatkan data mengenai post dari user tersebut melalui API
