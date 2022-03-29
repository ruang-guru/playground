const { expect, describe, it } = require('@jest/globals');
const gradingStudents = require('./main')

describe("Grading Students", () => {
  it("should return the grades after rounding as appropriate", () => {
    let case1 = gradingStudents([73, 67, 38, 33])

    expect(case1).toEqual([75, 67, 40, 33])
  })
})