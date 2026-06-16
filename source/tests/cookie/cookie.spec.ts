// CookieStore API support
// Chromium ✅
// Edge ✅
// Firefox ❌
// Webkit ❌

import { test, expect } from '@playwright/test';

test.use({ browserName: "chromium" })

test('test cookieStore', async ({ page }) => {
  const consoleLogs: string[] = []
  page.on('console', msg => consoleLogs.push(msg.text()))

  await page.goto('/tests/cookie/public');

  // Wait a bit to collect logs
  await page.waitForTimeout(1000);

  expect(consoleLogs.some(log => log.includes("panic"))).toBeFalsy();
});
