package handlers

import (
	"context"
	"fmt"
	"link-crawler/helpers"
	"log"

	"github.com/chromedp/chromedp"
)

func ScrapeDynamicPage(url string, selectors map[string]interface{}) {
	ctx, cancel := chromedp.NewContext(context.Background())
	defer cancel()

	var htmlContent string
	err := chromedp.Run(ctx,
		chromedp.Navigate(url),
		chromedp.WaitVisible(selectors["product_container"].(string)),
		chromedp.OuterHTML("html", &htmlContent),
	)
	if err != nil {
		log.Fatal("Error scraping dynamic page:", err)
	}

	fmt.Println("Page content loaded. Now parsing...")

	// Parse the HTML content using GoQuery (delegating to a helper function)
	helpers.ParseProducts(htmlContent, selectors)
}
