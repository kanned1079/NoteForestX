package models

import (
	"time"

	"gorm.io/gorm"
)

// 文章状态
type ArticleStatus string

const (
	Draft     ArticleStatus = "draft"
	Published ArticleStatus = "published"
	Hidden    ArticleStatus = "hidden"
)

// ArticleTag 标签表
type ArticleTag struct {
	ID        string    `gorm:"type:char(36);primaryKey" json:"id"` // UUID
	Name      string    `gorm:"size:50;not null;unique" json:"name"`
	Code      *string   `gorm:"size:50;unique;default:null" json:"code"` // 可选稳定标识
	CreatedAt time.Time `json:"created_at"`
}

// Article 文章表
type Article struct {
	ID      string        `gorm:"type:char(36);primaryKey" json:"id"`
	Slug    *string       `gorm:"size:255;unique;default:null" json:"slug"`
	Title   string        `gorm:"size:255;not null" json:"title"`
	Top     bool          `gorm:"default:false" json:"top"`
	Status  ArticleStatus `gorm:"type:enum('draft','published','hidden');default:'draft'" json:"status"`
	Content string        `gorm:"type:longtext;not null" json:"content"`

	Tags []*ArticleTag `gorm:"many2many:x_article_tag_rel;joinForeignKey:ArticleID;joinReferences:TagID" json:"tags"`

	ImageUrl  *string        `json:"image_url"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deleted_at"`
}

// ArticleTagRel 文章-标签 关联表（中间表）
type ArticleTagRel struct {
	ArticleID string `gorm:"type:char(36);primaryKey" json:"article_id"`
	TagID     string `gorm:"type:char(36);primaryKey" json:"tag_id"`

	CreatedAt time.Time
}

func (ArticleTagRel) TableName() string {
	return "x_article_tag_rel"
}

func (ArticleTag) TableName() string {
	return "x_article_tag"
}

func (Article) TableName() string {
	return "x_article"
}
