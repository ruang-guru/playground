const { test, expect } = require("@playwright/test");
const path = require("path");

test("basic test", async ({ page }) => {
  await page.goto(
    `file:${path.join(
      __dirname,
      "../../../../frontend/dom/guess-number-assignment/index.html"
    )}`
  );

  page.locator("");

  const title = await page.locator(".input").type("1");

  await page.locator(".check").click();

  if (page.locator(".input") < page.locator("#hidden-number")) {
    page.locator(".message").should("have.text", "Too small, buddy!");
  }
});
