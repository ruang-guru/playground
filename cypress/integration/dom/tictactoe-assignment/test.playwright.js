const { test, expect } = require("@playwright/test");
const path = require("path");

test.beforeEach(async ({ page }) => {
  await page.goto(
    `file:${path.join(
      __dirname,
      "../../../../frontend/dom/tictactoe-assignment/index.html"
    )}`
  );
});

const { describe } = test;

describe("tictactoe", () => {
  describe("board", () => {
    test("initially has 3x3 buttons with empty text", async ({ page }) => {
      expect(await page.locator("[id='0-0']").textContent()).toBe("");
      expect(await page.locator("[id='0-1']").textContent()).toBe("");
      expect(await page.locator("[id='0-2']").textContent()).toBe("");
      expect(await page.locator("[id='1-0']").textContent()).toBe("");
      expect(await page.locator("[id='1-1']").textContent()).toBe("");
      expect(await page.locator("[id='1-2']").textContent()).toBe("");
      expect(await page.locator("[id='2-0']").textContent()).toBe("");
      expect(await page.locator("[id='2-1']").textContent()).toBe("");
      expect(await page.locator("[id='2-2']").textContent()).toBe("");
    });

    test("change the button text content when clicked", async ({ page }) => {
      await page.locator("[id='0-0']").click();
      expect(await page.locator("[id='0-0']").textContent()).toBe("X");

      await page.locator("[id='0-1']").click();
      expect(await page.locator("[id='0-1']").textContent()).toBe("O");
    });

    describe("when a button has been clicked", () => {
      test("will be disabled", async ({ page }) => {
        await page.locator("[id='0-0']").click();
        expect(
          await page.locator("[id='0-0']").getAttribute("disabled")
        ).not.toBe(null);

        await page.locator("[id='0-1']").click();
        expect(
          await page.locator("[id='0-1']").getAttribute("disabled")
        ).not.toBe(null);
      });
    });

    describe("when the winner has been decided", () => {
      test("reset the board", async ({ page }) => {
        await page.locator("[id='0-0']").click();
        await page.locator("[id='1-0']").click();
        await page.locator("[id='1-1']").click();
        await page.locator("[id='0-1']").click();
        await page.locator("[id='2-2']").click();

        expect(await page.locator("[id='0-0']").textContent()).toBe("");
        expect(await page.locator("[id='0-1']").textContent()).toBe("");
        expect(await page.locator("[id='0-2']").textContent()).toBe("");
        expect(await page.locator("[id='1-0']").textContent()).toBe("");
        expect(await page.locator("[id='1-1']").textContent()).toBe("");
        expect(await page.locator("[id='1-2']").textContent()).toBe("");
        expect(await page.locator("[id='2-0']").textContent()).toBe("");
        expect(await page.locator("[id='2-1']").textContent()).toBe("");
        expect(await page.locator("[id='2-2']").textContent()).toBe("");
      });
    });

    describe("when all buttons have been clicked without a winner", () => {
      test("reset the board", async ({ page }) => {
        // X O X
        // X X O
        // O X O
        await page.locator("[id='0-0']").click();
        await page.locator("[id='0-1']").click();
        await page.locator("[id='1-0']").click();
        await page.locator("[id='2-0']").click();
        await page.locator("[id='1-1']").click();
        await page.locator("[id='2-2']").click();
        await page.locator("[id='2-1']").click();
        await page.locator("[id='1-2']").click();
        await page.locator("[id='0-2']").click();

        expect(await page.locator("[id='0-0']").textContent()).toBe("");
        expect(await page.locator("[id='0-1']").textContent()).toBe("");
        expect(await page.locator("[id='0-2']").textContent()).toBe("");
        expect(await page.locator("[id='1-0']").textContent()).toBe("");
        expect(await page.locator("[id='1-1']").textContent()).toBe("");
        expect(await page.locator("[id='1-2']").textContent()).toBe("");
        expect(await page.locator("[id='2-0']").textContent()).toBe("");
        expect(await page.locator("[id='2-1']").textContent()).toBe("");
        expect(await page.locator("[id='2-2']").textContent()).toBe("");
      });
    });

    describe("scoreboard", () => {
      test("initially has zero scores", async ({ page }) => {
        expect(await page.locator("#o-wins").textContent()).toBe("0");
        expect(await page.locator("#x-wins").textContent()).toBe("0");
      });

      test("increments X's score when X wins", async ({ page }) => {
        await page.locator("[id='0-0']").click();
        await page.locator("[id='1-0']").click();
        await page.locator("[id='1-1']").click();
        await page.locator("[id='0-1']").click();
        await page.locator("[id='2-2']").click();

        expect(await page.locator("#x-wins").textContent()).toBe("1");
        expect(await page.locator("#o-wins").textContent()).toBe("0");
      });

      test("increments O's score when O wins", async ({ page }) => {
        await page.locator("[id='0-0']").click();
        await page.locator("[id='1-0']").click();
        await page.locator("[id='2-2']").click();
        await page.locator("[id='1-1']").click();
        await page.locator("[id='2-0']").click();
        await page.locator("[id='1-2']").click();

        expect(await page.locator("#x-wins").textContent()).toBe("0");
        expect(await page.locator("#o-wins").textContent()).toBe("1");
      });
    });
  });
});
