# URL Shortener

URL shortener merupakan sebuah aplikasi yang dapat memangkas URL yang panjang dengan membuat URL baru yang lebih pendek.

Buatlah sebuah HTTP server yang berfungsi sebagai URL shortener dengan requirement sebagai berikut:

- Ketika server menerima request URL, server harus dapat membuat alias / short url dari URL yang di request.
- Ketika client membuka short url yang diberikan oleh server, maka server harus redirect ke URL.
- Client dapat membuat custom short URL.

Gunakan database dengan menggunakan map.

Pada assignment ini terdapat 4 folder yaitu

- entity

    Digunakan untuk menyimpan model dari objek yang akan digunakan

- handlers

    Digunakan untuk memanajemen handler.

- repository

    Digunakan untuk implementasi data layer.

- router

    Digunakan untuk memanajemen route.
