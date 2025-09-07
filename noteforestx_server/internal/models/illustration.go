package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"noteforestx_server/internal/config"
	"time"
)

// 插画表
type Illustration struct {
	Id        uuid.UUID          `json:"id" gorm:"type:char(36);primaryKey"`
	PixivId   string             `json:"pixiv_id" gorm:"uniqueIndex;size:100"` // PixivId 全局唯一
	FilePath  string             `json:"file_path"`
	Name      string             `json:"name"`      // 插画名
	AuthorId  uuid.UUID          `json:"author_id"` // 外键
	Author    IllustrationAuthor `json:"author" gorm:"foreignKey:AuthorId;references:Id"`
	Link      string             `json:"link"` // 插画原链接
	Tags      []IllustrationTag  `json:"tags" gorm:"many2many:illustration_tag_mapping;"`
	CreatedAt *time.Time         `json:"created_at"`
	UpdatedAt *time.Time         `json:"updated_at"`
	DeletedAt gorm.DeletedAt     `json:"deleted_at"`
}

// 插画作者表
type IllustrationAuthor struct {
	Id        uuid.UUID      `json:"id" gorm:"type:char(36);primaryKey"`
	Name      string         `json:"name" gorm:"size:255;index"`
	Link      string         `json:"link" gorm:"size:512"`
	CreatedAt *time.Time     `json:"created_at"`
	UpdatedAt *time.Time     `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at"`
}

// 插画标签表
type IllustrationTag struct {
	Id        uuid.UUID      `json:"id" gorm:"type:char(36);primaryKey"`
	Name      string         `json:"name" gorm:"type:varchar(255);uniqueIndex"`
	CreatedAt *time.Time     `json:"created_at"`
	UpdatedAt *time.Time     `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at"`
}

// 插画-标签关系表 (多对多)
type IllustrationTagMapping struct {
	IllustrationId uuid.UUID      `json:"illustration_id" gorm:"type:char(36);primaryKey"`
	TagId          uuid.UUID      `json:"tag_id" gorm:"type:char(36);primaryKey"`
	CreatedAt      *time.Time     `json:"created_at"`
	UpdatedAt      *time.Time     `json:"updated_at"`
	DeletedAt      gorm.DeletedAt `json:"deleted_at"`
}

// 表名定义
func (Illustration) TableName() string {
	return config.ExistingAppConfig.DbConfig.TablePrefix + "illustration"
}

func (IllustrationAuthor) TableName() string {
	return config.ExistingAppConfig.DbConfig.TablePrefix + "illustration_author"
}

func (IllustrationTag) TableName() string {
	return config.ExistingAppConfig.DbConfig.TablePrefix + "illustration_tag"
}

func (IllustrationTagMapping) TableName() string {
	return config.ExistingAppConfig.DbConfig.TablePrefix + "illustration_tag_mapping"
}
