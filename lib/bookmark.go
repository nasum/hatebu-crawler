package lib

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/PuerkitoBio/goquery"
	"github.com/gocolly/colly/v2"
)

func GetBookmark(c *colly.Collector) {
	cmd := flag.NewFlagSet("bookmark", flag.ExitOnError)
	target := cmd.String("target", "", "crawl target")
	err := cmd.Parse(os.Args[2:])

	if err != nil {
		log.Fatal(err)
	}

	bookmarkUrl := fmt.Sprintf("https://b.hatena.ne.jp/%s/bookmark", *target)
	fmt.Println(bookmarkUrl)

	c.OnHTML(".bookmark-item", func(e *colly.HTMLElement) {
		link := e.DOM.Find(".centerarticle-entry-title a[href]")
		title := link.Text()
		url, _ := link.Attr("href")
		fmt.Printf("%s: %s\n", title, url)

		var tags []string
		e.DOM.Find(".centerarticle-reaction-tags li").Each(func(_ int, s *goquery.Selection) {
			tags = append(tags, s.Find("a").Text())
		})

		fmt.Printf("tags: %s\n", tags)

		createdAt := e.DOM.Find(".centerarticle-reaction-timestamp").Text()

		fmt.Printf("created_at: %s\n", createdAt)

	})

	err = c.Visit(bookmarkUrl)

	if err != nil {
		fmt.Printf("err: %s", err)
	}
}
