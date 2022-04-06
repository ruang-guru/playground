const { test, expect } = require("@playwright/test");
const path = require("path");

const { describe, beforeEach } = test;

beforeEach(async ({ page }) => {
  await page.goto(
    `file:${path.join(
      __dirname,
      "../../../../frontend/dom/counter-app/index.html"
    )}`
  );
});

describe("counter app", () => {
  describe("add", () => {
    test("adds the number", async ({ page }) => {
      await page.locator("#add").click();
      expect(await page.locator("#number").textContent()).toBe("1");

      await page.locator("#add").click();
      await page.locator("#add").click();
      expect(await page.locator("#number").textContent()).toBe("3");
    });
  });

  describe("number", () => {
    test("is initialized to 0", async ({ page }) => {
      expect(await page.locator("#number").textContent()).toBe("0");
    });
  });

  describe("substract", async () => {
    test("substract the number", async ({ page }) => {
      await page.locator("#add").click();
      await page.locator("#add").click();
      await page.locator("#add").click();
      await page.locator("#add").click();
      await page.locator("#add").click();
      await page.locator("#add").click();
      await page.locator("#add").click();
      await page.locator("#add").click();

      await page.locator("#substract").click();
      expect(await page.locator("#number").textContent()).toBe("7");

      await page.locator("#substract").click();
      await page.locator("#substract").click();

      expect(await page.locator("#number").textContent()).toBe("5");
    });

    test("rejects negative numbers", async ({ page }) => {
      await page.locator("#substract").click();

      expect(await page.locator("#message").textContent()).toBe(
        "Oops! you reach the min value!"
      );
    });
  });
});
