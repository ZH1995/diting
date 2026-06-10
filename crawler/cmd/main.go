package main

import (
	"github.com/ZH1995/diting/crawler/internal/collector"
	"github.com/ZH1995/diting/crawler/internal/spider"
)

func main() {
	spiders := []spider.Spider{
		&spider.SiteASpider{},
	}
	collector.Run(spiders)
}
