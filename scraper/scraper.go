package scraper

import (
	"fmt"

	"github.com/gocolly/colly"
)

func ScrapeInstructions(url string) []string {
	collector := colly.NewCollector()

	collector.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL)
	})
	collector.OnResponse(func(r *colly.Response) {
		fmt.Println("Got a response from", r.Request.URL)
	})
	collector.OnError(func(r *colly.Response, e error) {
		fmt.Println("Error:", e)
	})

	var instructions Instructions

	collector.OnHTML(`section#section--instructions_1-0`, func(e *colly.HTMLElement) {
		e.ForEach(`li`, func(i int, e *colly.HTMLElement) {
			instructions.url = url
			if e.ChildText("p.comp") != "" {
				instructions.steps = append(instructions.steps, e.ChildText("p.comp"))
			}
		})
	})

	collector.Visit(url)
	return instructions.steps
}

type Instructions struct {
	url   string
	steps []string
}
