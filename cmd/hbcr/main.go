package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/gocolly/colly/v2"
)

func main() {
	c := colly.NewCollector()

	if len(os.Args) < 2 {
		os.Exit(0)
	}

	switch os.Args[1] {
	case "user":
		cmd := flag.NewFlagSet("user", flag.ExitOnError)
		target := cmd.String("target", "", "crawl target")
		err := cmd.Parse(os.Args[2:])

		if err != nil {
			log.Fatal(err)
		}

		bookmarkUrl := fmt.Sprintf("https://b.hatena.ne.jp/%s/bookmark", *target)
		fmt.Println(bookmarkUrl)

		c.OnHTML("a[href]", func(e *colly.HTMLElement) {
			link := e.Attr("href")
			fmt.Printf("Link: %s\n", link)
		})

		err = c.Visit(bookmarkUrl)

		if err != nil {
			fmt.Printf("err: %s", err)
		}
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
