const { test, expect } = require("@playwright/test");
const path = require("path");

test.beforeEach(async ({ page }) => {
  await page.goto(
    `file:${path.join(
      __dirname,
      "../../../../frontend/dom/stopwatch-app-assignment/index.html"
    )}`
  );
});

const { describe } = test;

describe("start", () => {
  test("start the stopwatch", async ({ page }) => {
    expect(await page.locator("#stopwatch").textContent()).toBe("00:00:00");

    await page.locator("#start").click();
    await page.waitForTimeout(1000);

    expect(await page.locator("#stopwatch")).not.toBe("00:00:00");
  });
});

describe("stop", () => {
  test("stop the stopwatch", async ({ page }) => {
    expect(await page.locator("#stopwatch").textContent()).toBe("00:00:00");

    await page.locator("#start").click();
    await page.waitForTimeout(1000);

    await page.locator("#stop").click();

    const prevValue = await page.locator("#stopwatch").textContent();

    await page.waitForTimeout(1000);

    expect(await page.locator("#stopwatch").textContent()).toBe(prevValue);
  });
});

describe("reset", () => {
  test("reset the stopwatch", async ({ page }) => {
    expect(await page.locator("#stopwatch").textContent()).toBe("00:00:00");

    await page.locator("#start").click();
    await page.waitForTimeout(1000);

    await page.locator("#stop").click();

    await page.locator("#reset").click();

    expect(await page.locator("#stopwatch").textContent()).toBe("00:00:00");
  });
});
