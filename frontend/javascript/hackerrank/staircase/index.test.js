const { expect, describe, it } = require('@jest/globals');
const staircase = require('./main')

describe("Staircase", () => {
  it("should print stairs based on the size contained in the input", () => {
    let case1 = staircase(10)

    expect(case1).toMatch(`         #
        ##
       ###
      ####
     #####
    ######
   #######
  ########
 #########
##########`)
  })
})