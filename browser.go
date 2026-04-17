package playwrightbase

import (
	"context"
	"time"

	"github.com/playwright-community/playwright-go"
)

type Browser struct {
	pw         *playwright.Playwright
	browser    playwright.Browser
	activePage playwright.Page
	pages      []playwright.Page
}

func (b *Browser) InstallBrowser() error {
	err := playwright.Install()
	if err != nil {
		return err
	}
	return nil
}

func (b *Browser) OpenBrowser(ctx context.Context, args map[string]interface{}) error {
	b.InstallBrowser()

	headless := false
	if val, ok := args["headless"]; ok {
		if headlessVal, ok := val.(bool); ok {
			headless = headlessVal
		}
	}
	if b.pw == nil {
		pwTmp, err := playwright.Run()
		pw, err := pwTmp.Chromium.Launch(playwright.BrowserTypeLaunchOptions{
			Headless: playwright.Bool(headless),
			Args:     []string{"--no-sandbox", "--disable-setuid-sandbox"},
		})
		if err != nil {
			return err
		}
		b.pw = pwTmp
		b.browser = pw

		context, err := b.browser.NewContext(playwright.BrowserNewContextOptions{
			IgnoreHttpsErrors: playwright.Bool(true),
		})

		page, err := context.NewPage()
		if err != nil {
			return err
		}
		b.activePage = page
		b.pages = append(b.pages, page)

	}
	return nil
}

func (b *Browser) CloseBrowser() error {
	if b.browser != nil {
		err := b.browser.Close()
		if err != nil {
			return err
		}
		b.browser = nil
	}
	if b.pw != nil {
		err := b.pw.Stop()
		if err != nil {
			return err
		}
		b.pw = nil
	}
	return nil
}

func (b *Browser) Sleep(ctx context.Context, duration time.Duration) error {
	time.Sleep(duration)
	return nil
}
