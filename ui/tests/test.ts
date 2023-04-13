import { chromium, expect, test } from '@playwright/test';

test('home page redirects to auth page', async ({ page }) => {
  await page.goto('/');
  await page.waitForURL('**/auth');
  expect(page.getByRole('link', { name: 'Log in' }));
  expect(page.getByRole('link', { name: 'Sign up' }));
});

test('diary page redirects to auth page', async ({ page }) => {
  await page.goto('/diary');
  await page.waitForURL('**/auth');
  expect(page.getByRole('link', { name: 'Log in' }));
  expect(page.getByRole('link', { name: 'Sign up' }));
});

test('search page redirects to auth page', async ({ page }) => {
  await page.goto('/search');
  await page.waitForURL('**/auth');
  expect(page.getByRole('link', { name: 'Log in' }));
  expect(page.getByRole('link', { name: 'Sign up' }));
});

test('account page redirects to auth page', async ({ page }) => {
  await page.goto('/account');
  await page.waitForURL('**/auth');
  expect(page.getByRole('link', { name: 'Log in' }));
  expect(page.getByRole('link', { name: 'Sign up' }));
});

test('account settings page redirects to auth page', async ({ page }) => {
  await page.goto('/account/settings');
  await page.waitForURL('**/auth');
  expect(page.getByRole('link', { name: 'Log in' }));
  expect(page.getByRole('link', { name: 'Sign up' }));
});

test('account goals page redirects to auth page', async ({ page }) => {
  await page.goto('/account/goals');
  await page.waitForURL('**/auth');
  expect(page.getByRole('link', { name: 'Log in' }));
  expect(page.getByRole('link', { name: 'Sign up' }));
});

test('invalid page shows not found page', async ({ page }) => {
  await page.goto('/invalid-page');
  await expect(page.getByRole('heading', { name: '404' })).toBeVisible();
  await expect(page.getByRole('link', { name: 'Back to home' })).toBeVisible();
});

test('largest contentful paint', async () => {
  const browser = await chromium.launch();
  const page = await browser.newPage();
  await page.goto('/');
  // https://www.checklyhq.com/learn/headless/basics-performance/
  const largestContentfulPaint = await page.evaluate(() => {
    return new Promise((resolve) => {
      new PerformanceObserver((l) => {
        const entries = l.getEntries();
        // the last entry is the largest contentful paint
        const largestPaintEntry = entries.at(-1);
        resolve(largestPaintEntry?.startTime);
      }).observe({
        type: 'largest-contentful-paint',
        buffered: true,
      });
    });
  });

  expect(parseFloat(largestContentfulPaint as string)).toBeLessThanOrEqual(2000);

  await browser.close();
});
