const { expect, describe, it } = require('@jest/globals');
const Stack = require('./stack')
const { IsValidParentheses } = require('./index')

describe("IsValidParentheses", () => {
    describe("when given empty string", () => {
        it("should return false", () => {
            let got = IsValidParentheses("")
            expect(got).toEqual(false)
        })
    })
    describe("when given correct parentheses sequence", () => {
        it("should return true", () => {
            let got = IsValidParentheses("()")
            expect(got).toEqual(true)
        })
        it("should return true", () => {
            let got = IsValidParentheses("(){}[]")
            expect(got).toEqual(true)
        })
    })
    describe("when given unbalanced parentheses sequence", () => {
        it("should return false", () => {
            let got = IsValidParentheses("{[()]")
            expect(got).toEqual(false)
        })
        it("should return false", () => {
            let got = IsValidParentheses("[()]}")
            expect(got).toEqual(false)
        })
    })
    describe("when incorrect parentheses sequence", () => {
        it("should return false", () => {
            let got = IsValidParentheses("[{)]")
            expect(got).toEqual(false)
        })
    })
    describe("when given only open parentheses sequence", () => {
        it("should return false", () => {
            let got = IsValidParentheses("({{")
            expect(got).toEqual(false)
        })
    })
    describe("when given only close parentheses sequence", () => {
        it("should return false", () => {
            let got = IsValidParentheses("]])")
            expect(got).toEqual(false)
        })
    })
})
