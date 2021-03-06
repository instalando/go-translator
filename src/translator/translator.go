package translator

import (
	"context"

	"github.com/chromedp/chromedp"
)

func Translate(text string, langFrom string, langTo string) (string, error) {
	ctx, cancel := chromedp.NewContext(context.Background())
	defer cancel()

	var res string

	err := navigateToTranslator(ctx, langFrom, langTo)
	if err != nil {
		return "", err
	}

	err = waitVisibleTextarea(ctx)
	if err != nil {
		return "", err
	}

	err = sendKeysToTranslateInTextarea(ctx, text)
	if err != nil {
		return "", err
	}

	res, err = getTextTranslated(ctx)
	if err != nil {
		return "", err
	}

	return res, nil
}

func navigateToTranslator(ctx context.Context, langFrom string, langTo string) error {
	return chromedp.Run(ctx, chromedp.Navigate("https://translate.google.com.br/?"+"sl="+langFrom+"&tl="+langTo))
}

func waitVisibleTextarea(ctx context.Context) error {
	return chromedp.Run(ctx, chromedp.WaitVisible(`textarea[class="er8xn"]`, chromedp.NodeVisible))
}

func sendKeysToTranslateInTextarea(ctx context.Context, text string) error {
	return chromedp.Run(ctx, chromedp.SendKeys(`textarea[class="er8xn"]`, text, chromedp.NodeVisible))
}

func getTextTranslated(ctx context.Context) (string, error) {
	var res string

	err := chromedp.Run(ctx, chromedp.Text(`span[class="JLqJ4b ChMk0b"]`, &res, chromedp.NodeVisible))
	if err != nil {
		return "", err
	}

	return res, nil
}
