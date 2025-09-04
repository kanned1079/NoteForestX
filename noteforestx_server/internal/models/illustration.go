package models

import (
	"github.com/google/uuid"
	"gorm.io/datatypes"
	"gorm.io/gorm"
	"noteforestx_server/internal/config"
	"time"
)

type Illustration struct {
	Id            uuid.UUID      `json:"id" gorm:"type:char(36);primaryKey"`
	PixivId       string         `json:"pixiv_id"` // 也就是上传文件的文件名去掉格式后缀
	TagsId        datatypes.JSON `json:"tags_id" gorm:"type:json"`
	FilePath      string         `json:"file_path"`
	Name          string         `json:"name"`            // 插画名
	Author        string         `json:"author"`          // 作者
	PixivAuthorId string         `json:"pixiv_author_id"` // pixiv作者id
	PixivLink     string         `json:"pixiv_link"`      // 插画原链接
	CreatedAt     *time.Time     `json:"created_at"`
	UpdatedAt     *time.Time     `json:"updated_at"`
	DeletedAt     gorm.DeletedAt `json:"deleted_at"`
}

func (Illustration) TableName() string {
	return config.ExistingAppConfig.DbConfig.TablePrefix + "illustration"
}
