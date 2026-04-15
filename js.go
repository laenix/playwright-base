package playwrightbase

import "context"

func (b *Browser) Evaluate(ctx context.Context, script string, args ...interface{}) (interface{}, error) {
	page, err := b.GetActivePage(ctx)
	if err != nil {
		return nil, err
	}
	result, err := page.Evaluate(script, args...)
	if err != nil {
		return nil, err
	}
	return result, nil
}
