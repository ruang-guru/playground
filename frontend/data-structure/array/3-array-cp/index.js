// Kalimat adalah daftar kata yang dipisahkan oleh satu spasi tanpa spasi awal atau akhir.
// Kita diberikan serangkaian string dengan nama sentences, di mana setiap sentences[i] mewakili satu kalimat.
// Kembalikan jumlah maksimum kata yang muncul dalam satu kalimat.
//
// Contoh 1
// Input: sentences = ["Andi suka bermain bola", "Saya sedang belajar struktur data", "Terima kasih"]
// Output: 5
// Explanation: jumlah kata terbanyak ada di "Saya sedang belajar struktur data", karena terdapat 5 kata dibandingkan kalimat lainnya, 4 dan 2 kata
//
// Contoh 2
// Input: sentences = ["Andi suka bermain bola", "Budi suka bermain basket", "Terima kasih"]
// Output: 4
// Explanation: jumlah kata terbanyak ada di "Andi suka bermain bola" dan "Budi suka bermain basket", karena terdapat 4 kata dibandingkan kalimat lainnya, 2 kata
// Ada kemungkinan bahwa beberapa kalimat mengandung jumlah kata yang sama.

function mostWordsFound(sentences) {
    let max = 0;
    for (let i = 0; i < sentences.length; i++) {
        let count = countWords(sentences[i]);
        // TODO: answer here
    }
    return max
}

function countWords(sentence) {
    return 0; // TODO: replace this
}

module.exports = {
    mostWordsFound, countWords
}