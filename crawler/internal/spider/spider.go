package spider

import (
	"github.com/ZH1995/diting/crawler/internal/model"
	"github.com/gocolly/colly/v2"
)

type Spider interface {
	Name() string
	Domain() string
	StartURLs() []string
	RegisterHandlers(c *colly.Collector)
	ParseResponse(c *colly.Collector, e *colly.Response) (*model.HotItem, error)
}
