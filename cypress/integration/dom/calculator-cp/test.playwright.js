const { test, expect } = require("@playwright/test");
const path = require("path");

const { describe, beforeEach } = test;

beforeEach(async ({ page }) => {
  await page.goto(
    `file:${path.join(
      __dirname,
      "../../../../frontend/dom/calculator-cp/index.html"
    )}`
  );
});

describe("calculator", () => {
  describe("numbers", () => {
    test("write the pressed number to the output text field", async ({
      page,
    }) => {
      await page.locator("#num1").click();
      await page.locator("#num3").click();
      expect(await page.locator("input").inputValue()).toBe("13");
    });
  });
});

describe("multiply", () => {
  test("multiply the numbers", async ({ page }) => {
    await page.locator("#num2").click();
    await page.locator("#mult").click();
    await page.locator("#num3").click();
    await page.locator("#calc").click();
    expect(await page.locator("input").inputValue()).toBe("6");
  });
});

describe("division", () => {
  test("divide the numbers", async ({ page }) => {
    await page.locator("#num6").click();
    await page.locator("#div").click();
    await page.locator("#num3").click();
    await page.locator("#calc").click();
    expect(await page.locator("input").inputValue()).toBe("2");
  });
});

describe("plus", () => {
  test("adds the numbers", async ({ page }) => {
    await page.locator("#num5").click();
    await page.locator("#plus").click();
    await page.locator("#num3").click();
    await page.locator("#calc").click();
    expect(await page.locator("input").inputValue()).toBe("8");
  });
});

describe("minus", () => {
  test("minus the numbers", async ({ page }) => {
    await page.locator("#num5").click();
    await page.locator("#min").click();
    await page.locator("#num3").click();
    await page.locator("#calc").click();
    expect(await page.locator("input").inputValue()).toBe("2");
  });
});

describe("dot", () => {
  test("adds decimal point", async ({ page }) => {
    await page.locator("#num1").click();
    await page.locator("#dot").click();
    await page.locator("#num3").click();

    await page.locator("#calc").click();
    expect(await page.locator("input").inputValue()).toBe("1.3");
  });

  test("allow decimal numbers operation", async ({ page }) => {
    await page.locator("#num1").click();
    await page.locator("#dot").click();
    await page.locator("#num3").click();

    await page.locator("#plus").click();

    await page.locator("#num2").click();
    await page.locator("#dot").click();
    await page.locator("#num5").click();

    await page.locator("#calc").click();

    expect(await page.locator("input").inputValue()).toBe("3.8");
  });
});

describe("del", () => {
  test("delete the previous input", async ({ page }) => {
    await page.locator("#num1").click();
    await page.locator("#num3").click();
    await page.locator("#del").click();

    expect(await page.locator("input").inputValue()).toBe("1");
  });
});

describe("ac", () => {
  test("clear the input", async ({ page }) => {
    await page.locator("#num1").click();
    await page.locator("#num3").click();
    await page.locator("#ac").click();

    expect(await page.locator("input").inputValue()).toBe("");
  });
});
