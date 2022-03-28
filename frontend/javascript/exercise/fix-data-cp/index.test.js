const fixData = require('./main.js')

describe('Fix Data', () => {
  it('Should eliminate all virus within the word with all the vowels grather than consonant', () => {
    const result1 = fixData('abc#ab#ueo')
    const expectedResult1 = `abcbabbueo`
    expect(result1).toMatch(expectedResult1)
  })

  it('Should eliminate all virus within the word with all the consonant grather than vowels', () => {
    const result2 = fixData('abc#ab#bcb')
    const expectedResult2 = `abcaababcb`
    expect(result2).toMatch(expectedResult2)
  })

  it('Should eliminate all virus within the word with all the consonant equal as vowels', () => {
    const result3 = fixData('aiu#bcd#ab')
    const expectedResult3 = `aiucbcdcab`
    expect(result3).toMatch(expectedResult3)
  })
})