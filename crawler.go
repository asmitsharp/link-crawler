package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/chromedp/chromedp"
)

type Product struct {
	name string
	link string
}

type Category struct {
}

type SubCategory struct {
}

func main() {
	// Create a context for the headless browser
	ctx, cancel := chromedp.NewContext(context.Background())
	defer cancel()

	// Timeout to avoid long waits
	ctx, cancel = context.WithTimeout(ctx, 15*time.Second)
	defer cancel()

	// Define variables to capture the scraped data
	var categories []string

	// Run the headless browser tasks
	err := chromedp.Run(ctx,
		// Step 1: Navigate to the Amazon Electronics page
		chromedp.Navigate("https://www.amazon.in/s?bbn=976419031&rh=n%3A976419031&dc&qid=1727351441&ref=lp_1389401031_ex_n_1"),

		// Step 2: Wait for the page to load completely
		chromedp.WaitVisible(`#nav-xshop-container`),

		// Step 3: Scrape the categories (adjust the selector if necessary)
		chromedp.Evaluate(`Array.from(document.querySelectorAll('li.a-spacing-micro')).map(el => el.innerText)`, &categories),
	)

	// Handle any errors
	if err != nil {
		log.Fatal(err)
	}

	// Print out the categories
	for i, category := range categories {
		fmt.Printf("Category %d: %s\n", i+1, category)
	}
}
