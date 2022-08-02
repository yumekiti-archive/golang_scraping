package main

import (
	"fmt"

	"github.com/gocolly/colly"
)

func main() {
	c := colly.NewCollector()

	c.OnHTML("body", func (e *colly.HTMLElement) {
		fmt.Println(e.Text)
	})

	// Start scraping on https://hackerspaces.org
	c.Visit("https://example.com/")
}