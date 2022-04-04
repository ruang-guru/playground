describe("guess-number", () => {
  describe("check-number", () => {
    it("check if input value smaller than random number", () => {
      cy.visit("frontend/dom/guess-number-assignment/index.html");

      cy.get(".input").type("1")

      cy.get(".check").click();

      if (cy.get(".input") < cy.get("#hidden-number")) {
        cy.get(".message").should("have.text", "Too small, buddy!");
      }
    })

    it("check if input value grether than random number", () => {
      cy.visit("frontend/dom/guess-number-assignment/index.html");

      cy.get(".input").type("10")

      cy.get(".check").click();

      if (cy.get(".input") > cy.get("#hidden-number")) {
        cy.get(".message").should("have.text", "Oww... that's too big!");
      }
    })

    it("check if input value smaller than 1", () => {
      cy.visit("frontend/dom/guess-number-assignment/index.html");

      cy.get(".input").type("0")

      cy.get(".check").click();

      cy.get(".message").should("have.text", "Guess any number between 1 and 10");
    })

    it("check if input value grether than 10", () => {
      cy.visit("frontend/dom/guess-number-assignment/index.html");

      cy.get(".input").type("11")

      cy.get(".check").click();

      cy.get(".message").should("have.text", "Guess any number between 1 and 10");
    })
  })

  describe("new-game", () => {
    it("reset message", () => {
      cy.get(".new-game").click()

      cy.get(".message").should("have.text", "Lets start guessing...")
    })
  })
})