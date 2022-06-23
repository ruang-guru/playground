// Kantin sekolah menawarkan sandwich berbentuk lingkaran dan persegi panjang saat istirahat makan siang, yang diwakilkan dengan angka 0 dan 1.
// Sandwich lingkaran = 0, dan sandwich persegi panjang = 1
// Semua siswa berdiri dalam antrian. Setiap siswa juga memiliki preferensi kesukaan terhadap sandwich berbentuk lingkaran atau persegi panjang
// Jumlah sandwich di kantin sama dengan jumlah siswa. Sandwich ditempatkan dalam tumpukan. Pada setiap langkah:
// 1. Jika siswa di depan antrian lebih suka sandwich paling atas dari tumpukan, mereka akan mengambilnya dan meninggalkan antrian.
// 2. Jika tidak, mereka akan meninggalkannya dan kembali antri ke ujung belakang antrian.
// 3. Ini berlanjut sampai tidak ada siswa yang mengantri ingin mengambil sandwich teratas dan dengan demikian tidak bisa makan.
// Preferensi siswa juga dilambangkan dengan angka 0 dan 1, siswa yang suka sandwich berbentuk lingkaran dilambangkan dengan 0, siswa yang suka sandwich berbentuk persegi panjang dilambangkan dengan 1.
// Diberikan dua array integer siswa dan sandwich di mana sandwich[i] adalah jenis sandwich ke-i dalam tumpukan (i = 0 adalah bagian atas tumpukan) dan students[j] adalah preferensi siswa ke-j dalam antrian awal (j = 0 adalah bagian depan antrian).
// Kembalikan jumlah siswa yang tidak bisa makan.
//
// Contoh 1
// Input: students = [1,1,0,0], sandwiches = [0,1,0,1]
// Output: 0
// Explanation:
// - Siswa paling depan meninggalkan sandwich paling atas dan kembali ke akhir baris sehingga students = [1,0,0,1].
// - Siswa paling depan meninggalkan sandwich paling atas dan kembali ke akhir baris sehingga students = [0,0,1,1].
// - Siswa depan mengambil sandwich paling atas dan meninggalkan barisan sehingga students = [0,1,1] dan sandwiches = [1,0,1].
// - Siswa paling depan meninggalkan sandwich paling atas dan kembali ke akhir baris membuat students = [1,1,0].
// - Siswa depan mengambil sandwich paling atas dan meninggalkan barisan sehingga students = [1,0] dan sandwiches = [0,1].
// - Siswa paling depan meninggalkan sandwich paling atas dan kembali ke akhir baris sehingga students = [0,1].
// - Siswa depan mengambil sandwich paling atas dan meninggalkan barisan membuat students = [1] dan sandwiches = [1].
// - Siswa depan mengambil sandwich paling atas dan meninggalkan barisan membuat students = [] dan sandwiches = [].
// Jadi semua siswa bisa makan.
//
// Contoh 2
// Input: students = [1,1,1,0,0,1], sandwiches = [1,0,0,0,1,1]
// Output: 3
// Explanation:
// - Siswa depan mengambil sandwich paling atas dan meninggalkan barisan sehingga students = [1,1,0,0,1] dan sandwiches = [0,0,0,1,1].
// - Siswa paling depan meninggalkan sandwich paling atas dan kembali ke akhir baris sehingga students = [1,0,0,1,1].
// - Siswa paling depan meninggalkan sandwich paling atas dan kembali ke akhir baris sehingga students = [0,0,1,1,1].
// - Siswa depan mengambil sandwich paling atas dan meninggalkan barisan sehingga students = [0,1,1,1] dan sandwiches = [0,0,1,1].
// - Siswa depan mengambil sandwich paling atas dan meninggalkan barisan sehingga students = [1,1,1] dan sandwiches = [0,1,1].
// - Siswa tersisa dengan preferensi 1, yaitu students = [1,1,1] dan sandwich paling atas adalah 0, sehingga siswa tidak dapat mengambil sandwich tersebut meskipun antri secara bergantian.
// Jadi semua siswa yang tidak bisa makan sandwich adalah 3 orang.

function countStudentsCantEat(students, sandwiches) {
	return 0 // TODO: replace this
}

console.log(countStudentsCantEat([1,1,1,0,0,1], [1,0,0,0,1,1])); // 3
console.log(countStudentsCantEat([1,1,0,0], [0,1,0,1])); // 0

module.exports = {
    countStudentsCantEat
}

