package spider

import (
	"time"

	"github.com/ZH1995/diting/crawler/internal/model"
	"github.com/gocolly/colly/v2"
)

// Spider接口
type Spider interface {
	Name() string
	Domain() string
	StartURLs() []string
	RegisterHandlers(c *colly.Collector)
	ParseResponse(e *colly.Response) ([]*model.HotItem, error)
	Config() SpiderConfig
}

// 爬虫配置结构体
type SpiderConfig struct {
	AllowedDomains []string      // 允许的域名（空=不限制）
	Parallelism    int           // 并发数
	Delay          time.Duration // 请求间隔
	RandomDelay    time.Duration // 随机延迟
	UserAgent      string        // User-Agent
	Timeout        time.Duration // HTTP超时
	CronExpr       string        // 调度cron表达式（如 "0 */30 * * * *" 每30分钟）
	Referer        string
}
