package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

type UserRole string

const (
	UserRoleAdmin UserRole = "ADMIN"
	UserRoleUser  UserRole = "USER"
)

type User struct {
	Id        uuid.UUID      `json:"id" gorm:"type:char(36);primaryKey"`
	Username  *string        `json:"username" gorm:"type:varchar(255)"`
	Email     string         `json:"email" gorm:"uniqueIndex;not null"`
	Role      UserRole       `json:"role" gorm:"type:enum('ADMIN','USER');default:'USER'"`
	Password  string         `json:"-" gorm:"not null"`
	CreatedAt *time.Time     `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt *time.Time     `json:"updated_at" gorm:"autoUpdateTime"`
	DeletedAt gorm.DeletedAt `json:"deleted_at" gorm:"index"`
}

func (User) TableName() string {
	return "x_n_user"
}
