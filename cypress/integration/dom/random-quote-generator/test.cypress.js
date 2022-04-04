describe("generate-random-quote", () => {
  describe("generate-quote", () => {
    it("display quote", () => {
      cy.visit("frontend/dom/random-quote-generator-cp/index.html");

      cy.get(".btn-generate").click()

      cy.get("#random-quote").should("not.be.empty")
    })

    it("display author", () => {
      cy.visit("frontend/dom/random-quote-generator-cp/index.html");

      cy.get(".btn-generate").click()

      cy.get(".author").should("not.be.empty")
    })
  })
})