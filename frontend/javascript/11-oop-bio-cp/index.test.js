const Person = require('./index')

describe("oop-bio", () => {
  describe("Getter", () => {
    it("should return name, age and job form Person", () => {
      let case1 = new Person("Jhon", 29, "Programmer")

      expect(case1.getName).toMatch("Jhon")
      expect(case1.getAge).toBe(29)
      expect(case1.getJob).toMatch("Programmer")
    })
  })

  describe("Setter", () => {
    it("should return name, age and job with new values", () => {
      let case1 = new Person("Jhon", 29, "Programmer")
      case1.setName = "Jane"
      case1.setAge = 25
      case1.setJob = "Engineer"

      expect(case1.getName).toMatch("Jane")
      expect(case1.getAge).toBe(25)
      expect(case1.getJob).toMatch("Engineer")
    })
  })
})