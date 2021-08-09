package lib

import (
	"fmt"
	"log"
	"strconv"

	"github.com/PuerkitoBio/goquery"
	"github.com/gocolly/colly/v2"
)

type BookmarkCrawler struct {
	Target    string
	Collector *colly.Collector
}

func (bmk *BookmarkCrawler) CreateURL() string {
	return fmt.Sprintf("https://b.hatena.ne.jp/%s/bookmark", bmk.Target)
}

func (bmk *BookmarkCrawler) GetEntries() error {

	bookmarkUrl := bmk.CreateURL()

	var bookmarkList BookMarkList

	bmk.Collector.OnHTML("li.bookmark-item", func(e *colly.HTMLElement) {
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

		bookmarkList = append(bookmarkList, bookmark)
	})

	bmk.Collector.OnScraped(func(_ *colly.Response) {
		err := bookmarkList.ShowJson()
		if err != nil {
			log.Fatal(err)
		}
	})

	err := bmk.Collector.Visit(bookmarkUrl)

	return err
}

func (bmk *BookmarkCrawler) GetBookmarkCount() error {
	bookmarkUrl := bmk.CreateURL()

	bmk.Collector.OnHTML(".userprofile-status-count", func(e *colly.HTMLElement) {
		count, err := strconv.Atoi(e.DOM.Text())

		if err != nil {
			log.Fatal(err)
		}

		fmt.Println(count)
	})

	err := bmk.Collector.Visit(bookmarkUrl)

	return err
}
