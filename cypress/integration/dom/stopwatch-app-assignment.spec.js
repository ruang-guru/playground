describe("stopwatch", () => {
    describe("start", () => {
        it("start the stopwatch", () => {
            cy.visit("frontend/dom/stopwatch-app-assignment/index.html");

            cy.get("#stopwatch").should("have.text", "00:00:00");

            cy.get("#start").click();
            cy.wait(1000);

            cy.get("#stopwatch").not("have.text", "00:00:00");
        })
    })

    describe("stop", () => {
        it("stop the stopwatch", () => {
            cy.visit("frontend/dom/stopwatch-app-assignment/index.html");

            cy.get("#stopwatch").should("have.text", "00:00:00");

            cy.get("#start").click();
            cy.wait(1000);

            cy.get("#stop").click();
            let prevValue = cy.get("#stopwatch").invoke("text").then(text => {
                cy.wait(1000);

                //it doesn't change
                cy.get("#stopwatch").should("have.text", text);
            })

        })
    })

    describe("reset", () => {
        it("reset the stopwatch", () => {
            cy.visit("frontend/dom/stopwatch-app-assignment/index.html");

            cy.get("#stopwatch").should("have.text", "00:00:00");

            cy.get("#start").click();
            cy.wait(1000);

            cy.get("#reset").click();
            cy.get("#stopwatch").should("have.text", "00:00:00");
        })
    })
})