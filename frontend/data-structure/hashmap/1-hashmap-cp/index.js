function ageDistribution(persons) {
    let result = new Map();
    // TODO: answer here
    return result
}

function groupByAge(persons) {
    let result = new Map();
    // TODO: answer here
    return result
}

let people = [
    { name: "Bob", age: 21 },
    { name: "Sam", age: 28 },
    { name: "Ann", age: 21 },
    { name: "Joe", age: 22 },
    { name: "Ben", age: 28 },
]

console.log(ageDistribution(people));
console.log(groupByAge(people));

module.exports = {
    ageDistribution, groupByAge
}