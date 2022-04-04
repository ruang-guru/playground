// calculator.spec.js created with Cypress
//
// Start writing your Cypress tests below!
// If you're unfamiliar with how Cypress works,
// check out the link below and learn how to write your first test:
// https://on.cypress.io/writing-first-test
/// <reference types="cypress"/>

describe("calculator", () => {
    describe("numbers", () => {
        it("write the pressed number to the output text field", () => {
            cy.visit("frontend/dom/calculator-cp/index.html")
            cy.get("#num1").click()
            cy.get("#num3").click()
            cy.get("input").should("have.value", "13")
        })
    })

    describe("multiply", () => {
        it("multiply the numbers", () => {
            cy.visit("frontend/dom/calculator-cp/index.html")
            cy.get("#num2").click()
            cy.get("#mult").click()
            cy.get("#num3").click()
            cy.get("#calc").click()
            cy.get("input").should("have.value", "6")
        })
    })

    describe("division", () => {
        it("divide the numbers", () => {
            cy.visit("frontend/dom/calculator-cp/index.html")
            cy.get("#num6").click()
            cy.get("#div").click()
            cy.get("#num3").click()
            cy.get("#calc").click()
            cy.get("input").should("have.value", "2")
        })
    })

    describe("plus", () => {
        it("adds the numbers", () => {
            cy.visit("frontend/dom/calculator-cp/index.html")
            cy.get("#num5").click()
            cy.get("#plus").click()
            cy.get("#num3").click()
            cy.get("#calc").click()
            cy.get("input").should("have.value", "8")
        })
    })

    describe("minus", () => {
        it("minus the numbers", () => {
            cy.visit("frontend/dom/calculator-cp/index.html")
            cy.get("#num5").click()
            cy.get("#min").click()
            cy.get("#num3").click()
            cy.get("#calc").click()
            cy.get("input").should("have.value", "2")
        })
    })

    describe("dot", () => {
        it("adds decimal point", () => {
            cy.visit("frontend/dom/calculator-cp/index.html")
            cy.get("#num1").click()
            cy.get("#dot").click()
            cy.get("#num3").click()

            cy.get("#calc").click()
            cy.get("input").should("have.value", "1.3")
        })

        it("allow decimal numbers operation", () => {
            cy.visit("frontend/dom/calculator-cp/index.html")
            cy.get("#num1").click()
            cy.get("#dot").click()
            cy.get("#num3").click()

            cy.get("#plus").click()

            cy.get("#num2").click()
            cy.get("#dot").click()
            cy.get("#num5").click()

            cy.get("#calc").click()
            cy.get("input").should("have.value", "3.8")
        })
    })

    describe("del", () => {
        it("delete the previous input", () => {
            cy.visit("frontend/dom/calculator-cp/index.html")
            cy.get("#num1").click()
            cy.get("#num3").click()
            cy.get("#del").click()
            cy.get("input").should("have.value", "1")
        })
    })

    describe("ac", () => {
        it("clear all input", () => {
            cy.visit("frontend/dom/calculator-cp/index.html")
            cy.get("#num1").click()
            cy.get("#num3").click()
            cy.get("#ac").click()
            cy.get("input").should("have.value", "")
        })
    })
})
