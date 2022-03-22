// Menggunakan switch case, petakan nilai ujian di bawah ini ke dalam Nilai A B C D E
// A = 91-100
// B = 71-90
// C = 61-70
// D = 51-60
// E = <= 50

// Masukan suatu angka
const score = parseInt(prompt("Masukan nilai: "));

// TODO: answer here

switch (true) {
    case score <= 100:
        alert(`Nilai A untuk Score ${score}`)
        break;
    case score >= 91:
        alert(`Nilai A untuk Score ${score}`)
        break;
    case score >= 71:
        alert(`Nilai B untuk Score ${score}`)
        break;
    case score >= 61:
        alert(`Nilai C untuk Score ${score}`)
        break;
    case score >= 51:
        alert(`Nilai D untuk Score ${score}`)
        break;
    case score < 51:
        alert(`Nilai E untuk Score ${score}`)
        break;
    default:
        alert(`Nilai belom dimasukkan, ${null}`)
        break;
}