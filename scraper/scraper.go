package main

import (
	"fmt"
	"os"
	"github.com/gocolly/colly"
)

func main(){
  args := os.Args
  url := args[1]
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
      instructions.steps = append(instructions.steps, e.ChildText("p.comp"))
    })
  })
  collector.Visit(url)
  fmt.Println(instructions.steps)
}

type Instructions struct {
  url string
  steps []string
}