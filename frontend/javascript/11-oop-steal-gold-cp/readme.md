# Steal the Gold!

Pada soal kali ini kita akan membuat sebuah permainan dimana setiap playernya dapat melakukan pencurian terhadap player lain.

kalian diminta untuk melengkapi class Player yang tersedia. Berikut ketentuannya:

1. Class tersebut memiliki properti name, gold, stealChange.
   - `gold`: properti default yang tidak perlu didapatkan dari instansiasi dengan default value 50.
   - `stealChange`:  properti default yang tidak perlu didapatkan dari instansiasi dengan default value 0.25. 

2. Terapat getter dan setter untuk gold dan stealChange.

3. Buat method `steal` yang menerima param Player (player lawan/ player lain). Lakukan steal randomizer berdasarkan `stealChance`.
   - Jika berhasil, player lawan kehilangan 5 gold dan gold player (sendiri) bertambah 5 gold dan kembalikan message "Berhasil mencuri 5 gold".
   - Jika gagal, kembalikan message "Gagal mencuri, coba lain kali"
   - Jika gold lawan lebih kecil dari 5 gold, steal gagal dengan message "Lawan terlalu miskin"