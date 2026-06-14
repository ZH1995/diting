package collector

import (
	"time"

	"github.com/ZH1995/diting/crawler/internal/spider"
	"github.com/gocolly/colly/v2"
)

func Run(spiders []spider.Spider) {
	for _, s := range spiders {
		cfg := mergeConfig(s.Config())
		c := NewCollector(cfg)
		s.RegisterHandlers(c)
		for _, url := range s.StartURLs() {
			c.Visit(url)
		}
		c.Wait()
	}
}

func NewCollector(cfg spider.SpiderConfig) *colly.Collector {
	opts := []colly.CollectorOption{colly.Async(true)}
	if cfg.UserAgent != "" {
		opts = append(opts, colly.UserAgent(cfg.UserAgent))
	}
	if len(cfg.AllowedDomains) > 0 {
		opts = append(opts, colly.AllowedDomains(cfg.AllowedDomains...))
	}
	c := colly.NewCollector(opts...)
	c.Limit(&colly.LimitRule{
		DomainGlob:  "*",
		Parallelism: cfg.Parallelism,
		Delay:       cfg.Delay,
		RandomDelay: cfg.RandomDelay,
	})
	if cfg.Timeout > 0 {
		c.SetRequestTimeout(cfg.Timeout)
	}
	// 设置Referer
	c.OnRequest(func(r *colly.Request) {
		if cfg.Referer != "" {
			r.Headers.Set("Referer", cfg.Referer)
		}
	})

	return c
}

func mergeConfig(custom spider.SpiderConfig) spider.SpiderConfig {
	cfg := spider.SpiderConfig{
		Parallelism: 2,
		Delay:       1 * time.Second,
		RandomDelay: 200 * time.Millisecond,
		UserAgent:   "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/148.0.0.0 Safari/537.36",
		Timeout:     10 * time.Second,
		Referer:     "",
	}
	if custom.Parallelism > 0 {
		cfg.Parallelism = custom.Parallelism
	}
	if custom.Delay > 0 {
		cfg.Delay = custom.Delay
	}
	if custom.RandomDelay > 0 {
		cfg.RandomDelay = custom.RandomDelay
	}
	if custom.UserAgent != "" {
		cfg.UserAgent = custom.UserAgent
	}
	if custom.Timeout > 0 {
		cfg.Timeout = custom.Timeout
	}
	if len(custom.AllowedDomains) > 0 {
		cfg.AllowedDomains = custom.AllowedDomains
	}
	if custom.CronExpr != "" {
		cfg.CronExpr = custom.CronExpr
	}
	if custom.Referer != "" {
		cfg.Referer = custom.Referer
	}
	return cfg
}
