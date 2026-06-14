package model

// 各平台来源标识
const (
	SourceBilibili uint8 = 1 // B站
)

/*

CREATE DATABASE IF NOT EXISTS diting DEFAULT CHARSET utf8mb4 DEFAULT COLLATE utf8mb4_unicode_ci;

CREATE TABLE IF NOT EXISTS hot_item (
	id             BIGINT UNSIGNED    PRIMARY KEY AUTO_INCREMENT    COMMENT '主键ID',
	title          VARCHAR(255)       NOT NULL DEFAULT ''           COMMENT '标题',
	url            VARCHAR(500)       NOT NULL DEFAULT ''           COMMENT '链接',
	item_rank      TINYINT UNSIGNED   NOT NULL DEFAULT 0            COMMENT '排名',
	content        VARCHAR(500)       NOT NULL DEFAULT ''           COMMENT '内容(前500字)',
	author         VARCHAR(100)       DEFAULT ''                    COMMENT '作者',
	publish_time   BIGINT             DEFAULT 0                     COMMENT '发布时间',
	source         TINYINT UNSIGNED   DEFAULT 0                     COMMENT '来源',
	title_hash     CHAR(32)           NOT NULL                      COMMENT '标题哈希值',
	create_time    BIGINT             DEFAULT 0                     COMMENT '创建时间',
	update_time    BIGINT             DEFAULT 0                     COMMENT '更新时间',
	UNIQUE KEY `uk_title_hash` (`title_hash`),
	KEY `idx_source` (`source`),
	KEY `idx_publish_time` (`publish_time`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT '热门条目表';


*/

type HotItem struct {
	ID          uint64 `gorm:"primaryKey;autoIncrement" json:"id"`
	Title       string `gorm:"type:varchar(255);not null;default:''" json:"title"`
	URL         string `gorm:"type:varchar(500);not null;default:''" json:"url"`
	ItemRank    uint8  `gorm:"type:tinyint unsigned;not null;default:0" json:"item_rank"`
	Content     string `gorm:"type:varchar(500);not null;default:''" json:"content"`
	Author      string `gorm:"type:varchar(100);default:''" json:"author"`
	PublishTime int64  `gorm:"type:bigint;default:0" json:"publish_time"`
	Source      uint8  `gorm:"type:tinyint unsigned;default:0" json:"source"`
	TitleHash   string `gorm:"type:char(32);not null;uniqueIndex:uk_title_hash" json:"title_hash"`
	CreateTime  int64  `gorm:"type:bigint;default:0" json:"create_time"`
	UpdateTime  int64  `gorm:"type:bigint;default:0" json:"update_time"`
}
