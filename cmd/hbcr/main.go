package main

import (
	"fmt"
	"os"

	"github.com/nasum/hatebu-crawler/lib"

	"github.com/gocolly/colly/v2"
)

func main() {
	c := colly.NewCollector()

	if len(os.Args) < 2 {
		os.Exit(0)
	}

	switch os.Args[1] {
	case "bookmark":
		lib.GetBookmark(c)
	case "top":
		c.OnHTML("a[href]", func(e *colly.HTMLElement) {
			link := e.Attr("href")
			fmt.Printf("Link: %s\n", link)
		})

		c.OnRequest(func(r *colly.Request) {
			fmt.Println("visitting", r.URL.String())
		})

		err := c.Visit("https://b.hatena.ne.jp/")

		if err != nil {
			fmt.Printf("err: %s", err)
		}
	}
}
