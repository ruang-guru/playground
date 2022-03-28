const { execSync } = require('child_process')
const fs = require('fs')

const reconstructedFilename = 'reconstructed.js'

const string = (word) => {
  
  let solution = fs.readFileSync(`${__dirname}/main.js`, "utf-8")

  solution = solution.replace(
    /(let|var) word .*/, `$1 word = ${typeof word === 'string' ? `"${word}"` : word}`
  )

  fs.writeFileSync(reconstructedFilename, solution)

  return String(execSync(`node ${reconstructedFilename}`))
}

afterAll(() => {
  if (fs.existsSync(reconstructedFilename)) {
    fs.unlinkSync(reconstructedFilename)
  }
})

describe('Breaking sentence', () => {
  it('Should print words one per line, with length of each word shown too', () => {
    const result = string("wow JavaScript is so cool and i will become frontend developer")
    const expectedResult = `First Word: wow, with length: 3
Second Word: JavaScript, with length: 10
Third Word: is, with length: 2
Fourth Word: so, with length: 2
Fifth Word: cool, with length: 4
Sixth Word: and, with length: 3
Seventh Word: i, with length: 1
Eighth Word: will, with length: 4
Ninth Word: become, with length: 6
Tenth Word: frontend, with length: 8
Eleventh Word: developer, with length: 9`
    expect(result).toMatch(expectedResult)
  })
})