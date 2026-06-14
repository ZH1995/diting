package main

import (
	"log"

	"github.com/ZH1995/diting/crawler/internal/collector"
	"github.com/ZH1995/diting/crawler/internal/pipeline"
	"github.com/ZH1995/diting/crawler/internal/spider"
)

func main() {
	// 创建pipeline
	dsn := "root:123456@tcp(127.0.0.1:3306)/diting?charset=utf8mb4&parseTime=True&loc=Local"
	p, err := pipeline.NewMySQLPipeline(dsn)
	if err != nil {
		log.Fatal(err)
	}
	defer p.Close()

	// 注册所有的spider
	spiders := []spider.Spider{
		spider.NewBilibiliSpider(p),
	}
	// 启动爬虫
	collector.Run(spiders)
}
