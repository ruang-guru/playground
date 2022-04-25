// Diberikan string "jewels" yang mewakili jenis batu yang merupakan permata, dan string "stones" yang mewakili batu yang dimiliki.
// Setiap karakter dalam "stones" adalah jenis batu yang dimiliki.
// Kita ingin tahu berapa banyak batu yang dimiliki yang juga termasuk permata.
// Huruf peka (case sensitive) terhadap huruf besar-kecil, jadi "a" dianggap sebagai jenis batu yang berbeda dari "A".
//
// Contoh 1
// Input: jewels = "aA", stones = "aAAbbbb"
// Output: 3
// Explanation: ada 3 permata di dalam variabel stones, yaitu a, A, dan A
//
// Contoh 2
// Input: jewels = "z", stones = "ZZ"
// Output: 0
// Explanation: karena case sensitive, tidak ada permata di dalam variabel stones

function numJewelsInStones(jewel, stones) {
    jewelMap = new Map();
    for (let i = 0; i < jewel.length; i++) {
        jewelMap.set(jewel[i], true);
    }

    let count = 0;
    for (let i = 0; i < stones.length; i++) {
        if (jewelMap.has(stones[i])) {
            count++;
        }
    }
    return count
}

let jewels = "aA";
let stones = "aAAbbbb";

console.log(numJewelsInStones(jewels, stones));

jewels = "z";
stones = "ZZ";
console.log(numJewelsInStones(jewels, stones));
