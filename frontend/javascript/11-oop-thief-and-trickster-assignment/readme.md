# Thief and Trickster

> Kalian tidak akan bisa mengerjakan study-case ini apabila belum menyelesaikan steal gold assignment

Setelah kalian menyeleseikan permainan apa soal sebelumnya. sekarang saatnya kita tambahkan tantangan agar permainan menjadi lebih menarik.

Terdapat class `Thief` dan `Trickster` yang merupakan extends dari class `Player`. kalian diminta untuk melengkapi class tersebut dengan ketentuan dibawah ini.

## Thief

1. Class `Thief` yang memiliki properti tambahan job dengan value thief, dan skill tambahan `robbedBlind` berupa function.

2. `robbedBlind` memiliki efek mengubah stealChance menjadi 75% dengan membayar 10 gold. 
   - Apabila gold player (sendiri) kurang dari 10 gold, kembalikan message "Aku terlalu miskin"

## Trickster

1. Class `Trickster` yang memiliki tambahan job dengan value thief, dan skill `distractionPurse` berupa function.

2. `distractionPurse` memiliki efek yaitu, ketika class ini kecurian uang, dia memiliki 25% chance untuk mencuri balik 10 gold.
   - Apabila total uang lawan kurang dari 10 gold, curi semua uang lawan. Kembalikan message "Berhasil mencuri balik semua uang lawan"
   - Apabila berhasil uang dan uang lawan lebih dari 10 gold, kembalikan message "Berhasil mencuri balik 10 gold"
   - Apabila gagal, kembalikan message "Gagal mencuri balik"

3. Tedapat getter dan setter untuk mengubah value `DistractionPurse` chance yang dimiliki

Copy semua logic steal yang sudah kamu tulis di study case sebelumnya dan tambahkan logic untuk mengecek job dari korban pencurian. Apabila korbannya adalah Trickster, maka jalankan skill `distractionPurse` tersebut.