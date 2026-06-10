package spider

import (
	"github.com/ZH1995/diting/crawler/internal/model"
	"github.com/gocolly/colly/v2"
)

type SiteASpider struct{}

func (s *SiteASpider) Name() string {
	return "site_a"
}

func (s *SiteASpider) Domain() string {
	return "example.com"
}

func (s *SiteASpider) StartURLs() []string {
	return []string{"https://example.com"}
}

func (s *SiteASpider) RegisterHandlers(c *colly.Collector) {
	c.OnHTML("div.article", func(e *colly.HTMLElement) {

	})
	c.OnScraped(func(c *colly.Response) {

	})
}

func (s *SiteASpider) ParseResponse(c *colly.Collector, e *colly.Response) (*model.Article, error) {
	return &model.Article{}, nil
}
