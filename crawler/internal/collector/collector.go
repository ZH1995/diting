package collector

import (
	"time"

	"github.com/ZH1995/diting/crawler/internal/spider"
	"github.com/gocolly/colly/v2"
)

func Run(spiders []spider.Spider) {
	for _, s := range spiders {
		c := NewCollector(s.Domain())
		s.RegisterHandlers(c)
		for _, url := range s.StartURLs() {
			c.Visit(url)
		}
		c.Wait()
	}
}

func NewCollector(domain string) *colly.Collector {
	c := colly.NewCollector(
		colly.AllowedDomains(domain),
		colly.Async(true),
		colly.UserAgent("..."),
	)
	c.Limit(&colly.LimitRule{
		DomainGlob:  "*",
		Parallelism: 2,
		Delay:       1 * time.Second,
	})
	//middleware.Attach(c)
	return c
}
