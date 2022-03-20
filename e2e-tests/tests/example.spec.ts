const { test, expect } = require('@playwright/test');

test('test basic flow', async ({ page, baseURL }) => {
  // Go to https://webapp-trelore.cloud.okteto.net/constructors
  await page.goto(baseURL + '/constructors');

  // Click text=Drivers Standings
  await page.click('text=Drivers Standings');
  expect(page.url()).toBe(baseURL + '/drivers');

  // Click text=Constructors Standings
  await page.click('text=Constructors Standings');
  expect(page.url()).toBe(baseURL + '/constructors');
})