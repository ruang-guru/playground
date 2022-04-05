const { test, expect } = require("@playwright/test");
const path = require("path");

test("basic test", async ({ page }) => {
  await page.goto(
    `file:${path.join(
      __dirname,
      "../../../../frontend/dom/guess-number-assignment/index.html"
    )}`
  );

  const title = await page.locator(".input").type("1");

  await page.locator(".check").click();

  if (
    (await page.locator(".input").inputValue()) <
    (await page.locator("#hidden-number").textContent())
  ) {
    expect(await page.locator(".message").textContent()).toBe(
      "Too small, buddy!"
    );
  }
});
