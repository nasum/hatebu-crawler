package main

import (
	"fmt"
	"github.com/gocolly/colly/v2"
)

func main() {
	c := colly.NewCollector()

	c.OnHTML("a[href]", func(e *colly.HTMLElement) {
		link := e.Attr("href")
		fmt.Println("Link: %q -> %s\n", e.Text, link)
	})

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("visitting", r.URL.String())
	})

	c.Visit("https://b.hatena.ne.jp/")
}
