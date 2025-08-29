package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

// Document 数据库表 n_document_dev
type Document struct {
	ID        uuid.UUID      `json:"id" gorm:"type:char(36);primaryKey"`
	Title     string         `json:"title" gorm:"type:varchar(255);not null;index"`
	Subtitle  *string        `json:"subtitle" gorm:"type:varchar(255)"`
	Category  *string        `json:"category" gorm:"type:varchar(255)"`
	Content   string         `json:"content" gorm:"type:longtext;not null"`
	Show      bool           `json:"show" gorm:"default:false"`
	ImageURL  *string        `json:"image_url" gorm:"type:varchar(255)"`
	CreatedAt *time.Time     `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt *time.Time     `json:"updated_at" gorm:"autoUpdateTime"`
	DeletedAt gorm.DeletedAt `json:"deleted_at" gorm:"index"`
}

// TableName 指定表名
func (Document) TableName() string {
	return "x_n_document"
}
