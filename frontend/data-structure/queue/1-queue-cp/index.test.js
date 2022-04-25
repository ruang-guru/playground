const { expect, describe, it } = require('@jest/globals');
const { countStudentsCantEat } = require('./index');

describe("countStudentsCantEat", () => {
    describe("when the student queue and sandwitches scores are the same", () => {
        it("all students will take sandwitches so that no students can't eat", () => {
            let students = [1,1,1,0,0,1];
            let sandwiches = [1,1,1,0,0,1];
            expect(countStudentsCantEat(students, sandwiches)).toBe(0);
        })
    })

    describe("when the student queue and sandwitches scores are different", () => {
        it("some students can't eat sandwitches", () => {
            let students = [1];
            let sandwiches = [0];
            expect(countStudentsCantEat(students, sandwiches)).toBe(1);

            students = [1, 1];
            sandwiches = [0, 1];
            expect(countStudentsCantEat(students, sandwiches)).toBe(2);

            students = [1, 1];
            sandwiches = [1, 0];
            expect(countStudentsCantEat(students, sandwiches)).toBe(1);

            students = [1,1,1,0,0,1];
            sandwiches = [1,0,0,0,1,1];
            expect(countStudentsCantEat(students, sandwiches)).toBe(3);
        });
    });
});
