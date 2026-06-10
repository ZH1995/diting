package model

/*

CREATE DATABASE IF NOT EXISTS diting;

CREATE TABLE IF NOT EXISTS article (
	id             BIGINT UNSIGNED    PRIMARY KEY AUTO_INCREMENT    COMMENT '主键ID',
	title          VARCHAR(255)       NOT NULL DEFAULT ''           COMMENT '标题',
	url            VARCHAR(500)       NOT NULL DEFAULT ''           COMMENT '链接',
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
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT '文章表';


*/

type Article struct {
	ID          uint64 `json:"id"`
	Title       string `json:"title"`
	URL         string `json:"url"`
	Content     string `json:"content"`
	Author      string `json:"author"`
	PublishTime int64  `json:"publish_time"`
	Source      uint8  `json:"source"`
	TitleHash   string `json:"title_hash"`
	CreateTime  int64  `json:"create_time"`
	UpdateTime  int64  `json:"update_time"`
}
