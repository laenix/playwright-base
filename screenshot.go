package playwrightbase

import (
	"context"

	"github.com/playwright-community/playwright-go"
)

func (b *Browser) Screenshot(ctx context.Context, path string) error {
	page, err := b.GetActivePage(ctx)
	if err != nil {
		return err
	}
	_, err = page.Screenshot(playwright.PageScreenshotOptions{
		Path: playwright.String(path),
	})
	if err != nil {
		return err
	}
	return nil
}
