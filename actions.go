package playwrightbase

import (
	"context"

	"github.com/playwright-community/playwright-go"
)

func (b *Browser) Click(ctx context.Context, selector string) error {
	page, err := b.GetActivePage(ctx)
	if err != nil {
		return err
	}
	err = page.Locator(selector).Click()
	if err != nil {
		return err
	}
	return nil
}

func (b *Browser) Fill(ctx context.Context, selector string, text string) error {
	page, err := b.GetActivePage(ctx)
	if err != nil {
		return err
	}
	err = page.Locator(selector).Fill(text)
	if err != nil {
		return err
	}
	return nil
}

func (b *Browser) SelectOption(ctx context.Context, selector string, value string) error {
	page, err := b.GetActivePage(ctx)
	if err != nil {
		return err
	}
	_, err = page.Locator(selector).SelectOption(playwright.SelectOptionValues{
		Values: playwright.StringSlice(value),
	})
	if err != nil {
		return err
	}
	return nil
}
