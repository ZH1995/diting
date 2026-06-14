// internal/pipeline/mysql_pipeline.go
package pipeline

import (
	"fmt"
	"time"

	"github.com/ZH1995/diting/crawler/internal/model"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"gorm.io/gorm/schema"
)

type MySQLPipeline struct {
	db *gorm.DB
}

func NewMySQLPipeline(dsn string) (*MySQLPipeline, error) {
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true, // 禁用表名复数
		},
	})
	if err != nil {
		return nil, fmt.Errorf("gorm open failed: %w", err)
	}

	sqlDB, _ := db.DB()
	sqlDB.SetMaxOpenConns(30)                 // 最大打开连接数
	sqlDB.SetMaxIdleConns(10)                 // 最大空闲连接数
	sqlDB.SetConnMaxLifetime(5 * time.Minute) // 连接最大复用时间
	return &MySQLPipeline{db: db}, nil
}

func (p *MySQLPipeline) Process(data *model.HotItem) error {
	//fmt.Printf("[pipeline] rank=%d title=%s\n", data.ItemRank, data.Title)
	// 标题哈希冲突时更新，实现幂等写入
	return p.db.Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: "title_hash"}},
		DoUpdates: clause.AssignmentColumns([]string{"title", "url", "item_rank", "content", "author", "publish_time", "update_time"}),
	}).Create(data).Error
}

func (p *MySQLPipeline) Close() error {
	sqlDB, err := p.db.DB()
	if err != nil {
		return err
	}
	return sqlDB.Close()
}
