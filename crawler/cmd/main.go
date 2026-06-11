package main

import (
	"github.com/ZH1995/diting/crawler/internal/collector"
	"github.com/ZH1995/diting/crawler/internal/spider"
)

func main() {
	// 注册所有的spider
	spiders := []spider.Spider{
		&spider.BilibiliSpider{},
	}
	// 启动爬虫
	collector.Run(spiders)
}
