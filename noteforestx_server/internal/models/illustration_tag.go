package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"noteforestx_server/internal/config"
	"time"
)

type IllustrationTag struct {
	Id        uuid.UUID      `json:"id" gorm:"type:char(36);primaryKey"`
	Name      string         `json:"name"  gorm:"type:varchar(255)"`
	CreatedAt *time.Time     `json:"created_at"`
	UpdatedAt *time.Time     `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at"`
}

func (IllustrationTag) TableName() string {
	return config.ExistingAppConfig.DbConfig.TablePrefix + "illustration_tag"
}
