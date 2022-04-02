const { test, expect, chromium } = require('@playwright/test');
test('test', async ({ page, baseURL }) => {

  // Go to http://localhost:8081/
  await page.goto(baseURL);

  // Click text=Constructors
  await page.click('text=Constructors');
  await page.waitForURL("**/constructors");
  expect(page.url()).toBe(baseURL + '/constructors');
  
  // Click text=Drivers
  await page.click('text=Drivers');
  await page.waitForURL("**/drivers");
  expect(page.url()).toBe(baseURL + '/drivers');

});

