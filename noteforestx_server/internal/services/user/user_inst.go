package user

import (
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
	"noteforestx_server/utils"
)

type UserService struct {
	utils utils.Utils
	Db    *gorm.DB
	Rdb   *redis.Client
}
