/// <reference types="cypress"/>

describe("counter-app", () => {
    describe("add", () => {
        it("adds the number", () => {
            cy.visit("frontend/dom/counter-app/index.html");

            cy.get("#add").click();
            cy.get("#number").should("have.text", "1");

            cy.get("#add").click();
            cy.get("#add").click();

            cy.get("#number").should("have.text", "3");
        })
    })

    describe("number", () => {
        it("is initialized to 0", () => {
            cy.visit("frontend/dom/counter-app/index.html");

            cy.get("#number").should("have.text", "0");
        })
    })

    describe("substract", () => {
        it("substract the number", () => {
            cy.visit("frontend/dom/counter-app/index.html");

            cy.get("#add").click();
            cy.get("#add").click();
            cy.get("#add").click();
            cy.get("#add").click();
            cy.get("#add").click();
            cy.get("#add").click();
            cy.get("#add").click();
            cy.get("#add").click();

            cy.get("#substract").click();
            cy.get("#number").should("have.text", "7");

            cy.get("#substract").click();
            cy.get("#substract").click();

            cy.get("#number").should("have.text", "5");
        })

        it("rejects negative number", () => {
            cy.visit("frontend/dom/counter-app/index.html");
            cy.get("#substract").click();
            cy.get("#message").should("have.text", "Oops! you reach the min value!");
        })
    })
})

