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
	ID        string  `gorm:"type:char(36);primaryKey"` // UUID
	Name      string  `gorm:"size:50;not null;unique"`
	Code      *string `gorm:"size:50;unique;default:null"` // 可选稳定标识
	CreatedAt time.Time
}

// Article 文章表
type Article struct {
	ID        string        `gorm:"type:char(36);primaryKey"`     // UUID
	Slug      *string       `gorm:"size:255;unique;default:null"` // 可选 slug
	Title     string        `gorm:"size:255;not null"`
	Top       bool          `gorm:"default:false"` // 是否置顶
	Status    ArticleStatus `gorm:"type:enum('draft','published','hidden');default:'draft'"`
	Content   string        `gorm:"type:longtext;not null"`
	Tags      []*ArticleTag `gorm:"many2many:article_tags;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"` // 软删除
}

func (ArticleTag) TableName() string {
	return "x_article_tag"
}

func (Article) TableName() string {
	return "x_article"
}
