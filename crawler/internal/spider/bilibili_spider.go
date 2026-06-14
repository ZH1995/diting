package spider

import (
	"encoding/json"
	"fmt"
	"log"
	"time"

	"github.com/ZH1995/diting/crawler/internal/model"
	"github.com/ZH1995/diting/crawler/internal/pipeline"
	"github.com/ZH1995/diting/crawler/internal/util"
	"github.com/gocolly/colly/v2"
)

type BilibiliSpider struct {
	pipeline pipeline.Pipeline
}

func NewBilibiliSpider(p pipeline.Pipeline) *BilibiliSpider {
	return &BilibiliSpider{pipeline: p}
}

func (s *BilibiliSpider) Name() string {
	return "bilibili"
}

func (s *BilibiliSpider) Domain() string {
	return "bilibili.com"
}

func (s *BilibiliSpider) StartURLs() []string {
	return []string{"https://api.bilibili.com/x/web-interface/popular?ps=20&pn=1"}
}

func (s *BilibiliSpider) RegisterHandlers(c *colly.Collector) {
	c.OnResponse(func(r *colly.Response) {
		// 解析数据
		items, err := s.ParseResponse(r)
		if err != nil {
			log.Println("parse error:", err)
			return
		}
		for _, item := range items {
			if err := s.pipeline.Process(item); err != nil {
				log.Printf("pipeline process error: %v", err)
			}
		}
	})
}

func (s *BilibiliSpider) ParseResponse(r *colly.Response) ([]*model.HotItem, error) {
	var resp bilibiliResponse
	if err := json.Unmarshal(r.Body, &resp); err != nil {
		return nil, fmt.Errorf("json unmarshal failed: %w", err)
	}
	if resp.Code != 0 {
		return nil, fmt.Errorf("api error: %s", resp.Message)
	}
	var items []*model.HotItem
	unixTime := time.Now().Unix()
	for i, item := range resp.Data.List {
		items = append(items, &model.HotItem{
			Title:       item.Title,
			URL:         item.ShortLinkV2,
			ItemRank:    uint8(i + 1),
			Content:     util.Truncate(util.StripHTML(item.Desc), 500),
			Author:      item.Owner.Name,
			PublishTime: int64(item.Pubdate),
			TitleHash:   util.MD5Hash(item.Title),
			Source:      model.SourceBilibili,
			CreateTime:  unixTime,
			UpdateTime:  unixTime,
		})
	}

	return items, nil
}

func (s *BilibiliSpider) Config() SpiderConfig {
	return SpiderConfig{
		AllowedDomains: []string{"bilibili.com", "api.bilibili.com"},
		Parallelism:    3,
		Delay:          500 * time.Millisecond,
		CronExpr:       "0 */30 * * * *",
		Referer:        "https://www.bilibili.com/v/popular/all",
	}
}

type bilibiliResponse struct {
	Code    int          `json:"code"`
	Message string       `json:"message"`
	Data    bilibiliData `json:"data"`
}

type bilibiliData struct {
	List []bilibiliItem `json:"list"`
}

type bilibiliItem struct {
	Title       string        `json:"title"`
	Pubdate     int           `json:"pubdate"`
	Desc        string        `json:"desc"`
	Owner       bilibiliOwner `json:"owner"`
	ShortLinkV2 string        `json:"short_link_v2"`
}

type bilibiliOwner struct {
	Name string `json:"name"`
}
