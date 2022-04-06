const { test, expect } = require("@playwright/test");
const path = require("path");

test.beforeEach(async ({ page }) => {
  await page.goto(
    `file:${path.join(
      __dirname,
      "../../../../frontend/dom/guess-number-assignment/index.html"
    )}`
  );
});

test("check if input value smaller than random number", async ({ page }) => {
  await page.locator(".input").type("1");

  await page.locator(".check").click();
  const hidden_number = Number(
    await page.locator("#hidden-number").textContent()
  );
  expect(hidden_number).toBeGreaterThan(0);

  if (1 < hidden_number) {
    expect(await page.locator(".message").textContent()).toBe(
      "Too small, buddy!"
    );
  }
});

test("check if input value greater than random number", async ({ page }) => {
  await page.locator(".input").type("10");

  await page.locator(".check").click();

  if (10 > (await page.locator("#hidden-number").textContent())) {
    expect(await page.locator(".message").textContent()).toBe(
      "Oww... that's too big!"
    );
  }
});

test("check if input value smaller than 1", async ({ page }) => {
  await page.locator(".input").type("0");

  await page.locator(".check").click();

  expect(await page.locator(".message").textContent()).toBe(
    "Guess any number between 1 and 10"
  );
});

test("check if input value greater than 10", async ({ page }) => {
  await page.locator(".input").type("11");

  await page.locator(".check").click();

  expect(await page.locator(".message").textContent()).toBe(
    "Guess any number between 1 and 10"
  );
});

test("reset message", async ({ page }) => {
  await page.locator(".new-game").click();

  expect(await page.locator(".message").textContent()).toBe(
    "Lets start guessing..."
  );
});
