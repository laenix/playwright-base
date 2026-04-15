package playwrightbase

import "context"

func (b *Browser) GetHTML(ctx context.Context, selector string) (string, error) {
	page, err := b.GetActivePage(ctx)
	if err != nil {
		return "", err
	}
	var content string
	if selector == "" {
		content, err = page.Content()
		if err != nil {
			return "", err
		}
	} else {
		content, err = page.Locator(selector).InnerHTML()
		if err != nil {
			return "", err
		}
	}
	return content, nil
}

func (b *Browser) GetTitle(ctx context.Context) (string, error) {
	page, err := b.GetActivePage(ctx)
	if err != nil {
		return "", err
	}
	title, err := page.Title()
	if err != nil {
		return "", err
	}
	return title, nil
}
