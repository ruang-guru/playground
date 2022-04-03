describe("tictactoe", () => {
    describe("board", () => {
        it("initially has 3x3 buttons with empty text", () => {
            cy.visit("frontend/dom/tictactoe-assignment/index.html");

            cy.get("button#0-0").should("have.text", "");
            cy.get("button#0-1").should("have.text", "");
            cy.get("button#0-2").should("have.text", "");
            cy.get("button#1-0").should("have.text", "");
            cy.get("button#1-1").should("have.text", "");
            cy.get("button#1-2").should("have.text", "");
            cy.get("button#2-0").should("have.text", "");
            cy.get("button#2-1").should("have.text", "");
            cy.get("button#2-2").should("have.text", "");
        })

        it("change the button text content when clicked", () => {
            cy.visit("frontend/dom/tictactoe-assignment/index.html");

            cy.get("button#0-0").click();
            cy.get("button#0-0").should("have.text", "X");

            cy.get("button#0-1").click();
            cy.get("button#0-1").should("have.text", "O");
        })

        context("when a button has been clicked", () => {
            it("will be disabled", () => {
                cy.visit("frontend/dom/tictactoe-assignment/index.html");

                cy.get("button#0-0").click();
                cy.get("button#0-0").should("have.attr", "disabled");

                cy.get("button#0-1").click();
                cy.get("button#0-1").should("have.attr", "disabled");
            })
        })

        context("when the winner has been decided", () => {
            it("reset the board", () => {
                cy.visit("frontend/dom/tictactoe-assignment/index.html");

                cy.get("button#0-0").click(); // X
                cy.get("button#1-0").click(); // O
                cy.get("button#1-1").click(); // X
                cy.get("button#0-1").click(); // O
                cy.get("button#2-2").click(); // X

                cy.get("button#0-0").should("have.text", "");
                cy.get("button#0-1").should("have.text", "");
                cy.get("button#0-2").should("have.text", "");
                cy.get("button#1-0").should("have.text", "");
                cy.get("button#1-1").should("have.text", "");
                cy.get("button#1-2").should("have.text", "");
                cy.get("button#2-0").should("have.text", "");
                cy.get("button#2-1").should("have.text", "");
                cy.get("button#2-2").should("have.text", "");
            })
        })

        context("when all buttons have been clicked without a winner", () => {
            it("reset the board", () => {
                cy.visit("frontend/dom/tictactoe-assignment/index.html");

                // X O X 
                // X X O
                // O X O 

                cy.get("button#0-0").click(); // X
                cy.get("button#0-1").click(); // O
                cy.get("button#1-0").click(); // X
                cy.get("button#2-0").click(); // O
                cy.get("button#1-1").click(); // X
                cy.get("button#2-2").click(); // O
                cy.get("button#2-1").click(); // X
                cy.get("button#1-2").click(); // O
                cy.get("button#0-2").click(); // X

                cy.get("button#0-0").should("have.text", "");
                cy.get("button#0-1").should("have.text", "");
                cy.get("button#0-2").should("have.text", "");
                cy.get("button#1-0").should("have.text", "");
                cy.get("button#1-1").should("have.text", "");
                cy.get("button#1-2").should("have.text", "");
                cy.get("button#2-0").should("have.text", "");
                cy.get("button#2-1").should("have.text", "");
                cy.get("button#2-2").should("have.text", "");

            })
        })
    })

    describe("scoreboard", () => {
        it("initially has zero scores", () => {
            cy.visit("frontend/dom/tictactoe-assignment/index.html");

            cy.get("#o-wins").should("have.text", "0");
            cy.get("#x-wins").should("have.text", "0");
        })

        it("increments X's score when X wins", () => {
            cy.visit("frontend/dom/tictactoe-assignment/index.html");

            cy.get("button#0-0").click(); // X
            cy.get("button#1-0").click(); // O
            cy.get("button#1-1").click(); // X
            cy.get("button#0-1").click(); // O
            cy.get("button#2-2").click(); // X

            cy.get("#x-wins").should("have.text", "1");
            cy.get("#o-wins").should("have.text", "0");
        })

        it("increments O's score when O wins", () => {
            cy.visit("frontend/dom/tictactoe-assignment/index.html");

            cy.get("button#0-0").click(); // X
            cy.get("button#1-0").click(); // O
            cy.get("button#2-2").click(); // X
            cy.get("button#1-1").click(); // O
            cy.get("button#2-0").click(); // X
            cy.get("button#1-2").click(); // O

            cy.get("#x-wins").should("have.text", "0");
            cy.get("#o-wins").should("have.text", "1");
        })
    })
})