package helpers

import (
	"fmt"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

func ParseProducts(html string, selectors map[string]interface{}) {
	doc, err := goquery.NewDocumentFromReader(strings.NewReader(html))
	if err != nil {
		fmt.Println("Error parsing HTML:", err)
		return
	}

	doc.Find(selectors["product_container"].(string)).Each(func(i int, s *goquery.Selection) {
		productLink, _ := s.Find(selectors["product_link"].(string)).Attr("href")
		productName := s.Find(selectors["product_link"].(string)).Text()
		fmt.Printf("Parsed product: %s (%s)\n", productName, productLink)
	})
}
