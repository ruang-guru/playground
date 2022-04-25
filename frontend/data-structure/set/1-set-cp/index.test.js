const { expect, describe, it } = require('@jest/globals');
const { intersection } = require('./index');

describe("intersection", () => {
    describe("when there is no intersection", () => {
        it("should return empty set", () => {
            let setA = new Set(["Java", "Python"]);
            let setB = new Set(["Swift"]);
            expect(intersection(setA, setB)).toEqual(new Set());
        });
    });

    describe("when there is intersection", () => {
        it("should return intersection set", () => {
            let setA = new Set(["Java", "Python", "Javascript", "C ++", "C#"]);
            let setB = new Set(["Java", "Python"]);
            expect(intersection(setA, setB)).toEqual(new Set(["Java", "Python"]));

            setA = new Set(["Agus", "Budi", "Charlie"]);
            setB = new Set(["Budi", "Agus"]);
            expect(intersection(setA, setB)).toEqual(new Set(["Agus", "Budi"]));
        }); 
    });
});
