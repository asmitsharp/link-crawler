package scraper

import (
	"fmt"
	"link-crawler/handlers"
)

func ScrapeSite(site string, config map[string]interface{}) {
	baseURL := config["base_url"].(string)
	categories := config["categories"].(map[string]interface{})
	selectors := config["selectors"].(map[string]interface{})

	for category, data := range categories {
		categoryURL := baseURL + data.(map[string]interface{})["url"].(string)
		fmt.Printf("Scraping %s category: %s\n", site, category)
		// Decide if the site needs Chromedp or Colly
		handlers.ScrapeDynamicPage(categoryURL, selectors)
	}
}
