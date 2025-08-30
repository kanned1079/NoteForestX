package models

import (
	"github.com/google/uuid"
	"noteforestx_server/internal/config"
)

type Illustration struct {
	Id uuid.UUID `json:"id" gorm:"type:char(36);primaryKey"`
}

func (Illustration) TableName() string {
	return config.ExistingAppConfig.DbConfig.TablePrefix + "illustration"
}
