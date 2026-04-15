package playwrightbase

import (
	"context"
	"fmt"

	"github.com/playwright-community/playwright-go"
)

func (b *Browser) NewPage(ctx context.Context) error {
	if b.browser == nil {
		return fmt.Errorf("no active browser page. call 'OpenBrowser' first")
	}
	page, err := b.browser.NewPage()
	if err != nil {
		return err
	}
	b.activePage = page
	b.pages = append(b.pages, page)

	return nil
}

func (b *Browser) Goto(ctx context.Context, url string) (string, error) {
	if b.activePage == nil {
		return "", fmt.Errorf("no active page. call 'NewPage' first")
	}
	_, err := b.activePage.Goto(url, playwright.PageGotoOptions{
		WaitUntil: playwright.WaitUntilStateLoad,
	})
	if err != nil {
		return "", err
	}
	return url, nil
}

func (b *Browser) GetUrl(ctx context.Context) (string, error) {
	if b.activePage == nil {
		return "", fmt.Errorf("no active page. call 'NewPage' first")
	}
	url := b.activePage.URL()
	return url, nil
}

func (b *Browser) WaitForSelector(ctx context.Context, selector string) error {
	if b.activePage == nil {
		return fmt.Errorf("no active page. call 'NewPage' first")
	}
	err := b.activePage.Locator(selector).WaitFor(playwright.LocatorWaitForOptions{
		State: playwright.WaitForSelectorStateAttached,
	})
	return err
}

func (b *Browser) GetActivePage(ctx context.Context) (playwright.Page, error) {
	if b.activePage == nil {
		return nil, fmt.Errorf("no active page. call 'NewPage' first")
	}
	return b.activePage, nil
}

func (b *Browser) GetPages(ctx context.Context) ([]playwright.Page, error) {
	if len(b.pages) == 0 {
		return nil, fmt.Errorf("no pages available. call 'NewPage' first")
	}
	return b.pages, nil
}

func (b *Browser) SwitchPage(ctx context.Context, index int) error {
	if index < 0 || index >= len(b.pages) {
		return fmt.Errorf("page index out of range")
	}
	b.activePage = b.pages[index]
	return nil
}

func (b *Browser) ClosePage(ctx context.Context, page playwright.Page) error {
	if page == nil {
		return fmt.Errorf("page is nil")
	}
	err := page.Close()
	if err != nil {
		return err
	}
	// Remove the closed page from the pages slice
	for i, p := range b.pages {
		if p == page {
			b.pages = append(b.pages[:i], b.pages[i+1:]...)
			break
		}
	}
	// If the closed page was the active page, reset activePage to nil
	if b.activePage == page {
		b.activePage = nil
	}
	return nil
}
