package lib

import (
	"fmt"
	"log"
	"strconv"

	"github.com/PuerkitoBio/goquery"
	"github.com/gocolly/colly/v2"
)

func CreateURL(target string) string {
	return fmt.Sprintf("https://b.hatena.ne.jp/%s/bookmark", target)
}

func GetEntries(target string) error {
	c := CreateCollector()
	bookmarkUrl := CreateURL(target)

	var bookmarkList BookMarkList

	c.OnHTML("li.bookmark-item", func(e *colly.HTMLElement) {
		link := e.DOM.Find(".centerarticle-entry-title a[href]")
		title := link.Text()
		url, _ := link.Attr("href")

		var tags []string
		e.DOM.Find(".centerarticle-reaction-tags li").Each(func(_ int, s *goquery.Selection) {
			tags = append(tags, s.Find("a").Text())
		})

		createdAt := e.DOM.Find(".centerarticle-reaction-timestamp").Text()

		bookmark := BookMark{
			Title:     title,
			URL:       url,
			Tags:      tags,
			CreatedAt: createdAt,
		}

		if bookmark.Title != "{{title}}" {
			bookmarkList = append(bookmarkList, bookmark)
		}
	})

	c.OnScraped(func(_ *colly.Response) {
		err := bookmarkList.ShowJson()
		if err != nil {
			log.Fatal(err)
		}
	})

	err := c.Visit(bookmarkUrl)

	return err
}

func GetTop() error {
	c := CreateCollector()

	c.OnHTML("a[href]", func(e *colly.HTMLElement) {
		link := e.Attr("href")
		fmt.Printf("Link: %s\n", link)
	})

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("visitting", r.URL.String())
	})

	err := c.Visit("https://b.hatena.ne.jp/")

	return err
}

func GetBookmarkCount(target string) error {
	c := CreateCollector()
	bookmarkUrl := CreateURL(target)

	c.OnHTML(".userprofile-status-count", func(e *colly.HTMLElement) {
		count, err := strconv.Atoi(e.DOM.Text())

		if err != nil {
			log.Fatal(err)
		}

		fmt.Println(count)
	})

	err := c.Visit(bookmarkUrl)

	return err
}
