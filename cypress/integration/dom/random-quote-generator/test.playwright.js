const { test, expect } = require("@playwright/test");
const path = require("path");

test.beforeEach(async ({ page }) => {
  await page.goto(
    `file:${path.join(
      __dirname,
      "../../../../frontend/dom/random-quote-generator-cp/index.html"
    )}`
  );
});

const { describe } = test;

describe("generate random quote", () => {
  describe("generate quote", () => {
    test("display quote", async ({ page }) => {
      await page.locator(".btn-generate").click();

      await expect(await page.locator("#random-quote")).not.toBeEmpty();
    });

    test("display author", async ({ page }) => {
      await page.locator(".btn-generate").click();

      await expect(await page.locator(".author")).not.toBeEmpty();
    });
  });
});
