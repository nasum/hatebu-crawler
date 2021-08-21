package lib

import (
	"github.com/gocolly/colly/v2"
)

func CreateCollector() *colly.Collector {
	return colly.NewCollector()
}
