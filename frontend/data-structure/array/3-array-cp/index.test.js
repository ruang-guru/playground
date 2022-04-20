const { expect, describe, it } = require('@jest/globals');
const { countWords, mostWordsFound } = require('./index');

describe("countWords", () => {
    describe("when input a sentence 2ith a few words", () => {
        it("returns the number of words", () => {
            let result = countWords("Andi suka bermain bola");
            expect(result).toBe(4);
        })
    })
})
describe("mostWordsFound", () => {
    describe("when input a list with each list having multiple sentences", () => {
        it("return number of maximum word count", () => {
            let result = mostWordsFound(["Andi suka bermain bola", "Saya sedang belajar struktur data", "Terima kasih"]);
            expect(result).toBe(5);
        })
    })
});
