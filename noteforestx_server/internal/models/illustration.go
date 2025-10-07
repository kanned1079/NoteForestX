package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// IllustrationImage 插畫圖片表 (一對多)
type IllustrationImage struct {
	Id             uuid.UUID      `json:"id" gorm:"type:char(36);primaryKey"`
	IllustrationId string         `json:"illustration_id" gorm:"size:100;index"` // 和 Illustration.Id 對應
	FilePath       string         `json:"file_path" gorm:"size:512"`
	Order          int            `json:"order"` // 在插畫中的順序，從0開始
	Width          int            `json:"width"`
	Height         int            `json:"height"`
	CreatedAt      *time.Time     `json:"created_at"`
	UpdatedAt      *time.Time     `json:"updated_at"`
	DeletedAt      gorm.DeletedAt `json:"deleted_at"`
}

// Illustration 插畫表 (插畫集)
type Illustration struct {
	Id          string              `json:"id" gorm:"size:100;primaryKey"`
	PixivId     string              `json:"pixiv_id" gorm:"size:100;uniqueIndex"`
	Name        string              `json:"name"`
	AuthorId    uuid.UUID           `json:"author_id"`
	Author      IllustrationAuthor  `json:"author" gorm:"foreignKey:AuthorId;references:Id"`
	Link        string              `json:"link"`
	Source      string              `json:"source" gorm:"size:50"`
	Description string              `json:"description" gorm:"type:text"`
	Tags        []IllustrationTag   `json:"tags" gorm:"many2many:x_illustration_tag_mapping;joinForeignKey:IllustrationId;joinReferences:TagId"`
	Images      []IllustrationImage `json:"images" gorm:"foreignKey:IllustrationId;references:Id"`
	Limited     bool                `json:"limited"`
	CreatedAt   *time.Time          `json:"created_at"`
	UpdatedAt   *time.Time          `json:"updated_at"`
	DeletedAt   gorm.DeletedAt      `json:"deleted_at"`
}

// IllustrationAuthor 插画作者表
type IllustrationAuthor struct {
	Id        uuid.UUID      `json:"id" gorm:"type:char(36);primaryKey"`
	Name      string         `json:"name" gorm:"size:255;index"`
	Link      string         `json:"link" gorm:"size:512"`
	CreatedAt *time.Time     `json:"created_at"`
	UpdatedAt *time.Time     `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at"`
}

// IllustrationTag 插画标签表
type IllustrationTag struct {
	Id   uuid.UUID `json:"id" gorm:"type:char(36);primaryKey"`
	Name string    `json:"name" gorm:"type:varchar(255);index"`
	//Name      string         `json:"name" gorm:"type:varchar(255);uniqueIndex"`
	CreatedAt *time.Time     `json:"created_at"`
	UpdatedAt *time.Time     `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at"`
}

// IllustrationTagMapping 插画-标签关系表 (多对多)
type IllustrationTagMapping struct {
	IllustrationId string         `json:"illustration_id" gorm:"size:100;primaryKey"`
	TagId          uuid.UUID      `json:"tag_id" gorm:"type:char(36);primaryKey"`
	CreatedAt      *time.Time     `json:"created_at"`
	UpdatedAt      *time.Time     `json:"updated_at"`
	DeletedAt      gorm.DeletedAt `json:"deleted_at"`
}

// 插画文件表名
func (IllustrationImage) TableName() string {
	return "x_illustration_image"
}

// 表名定义
func (Illustration) TableName() string {
	return "x_illustration"
}

func (IllustrationAuthor) TableName() string {
	return "x_illustration_author"
}

func (IllustrationTag) TableName() string {
	return "x_illustration_tag"
}

func (IllustrationTagMapping) TableName() string {
	return "x_illustration_tag_mapping"
}
