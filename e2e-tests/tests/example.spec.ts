const { test, expect } = require('@playwright/test');

test('test basic flow', async ({ page }) => {
  // Go to https://webapp-trelore.cloud.okteto.net/constructors
  await page.goto('https://webapp-trelore.cloud.okteto.net/constructors');

  // Click text=Drivers Standings
  await page.click('text=Drivers Standings');
  expect(page.url()).toBe('https://webapp-trelore.cloud.okteto.net/drivers');

  // Click text=Constructors Standings
  await page.click('text=Constructors Standings');
  expect(page.url()).toBe('https://webapp-trelore.cloud.okteto.net/constructors');
})