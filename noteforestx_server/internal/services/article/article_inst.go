package article

import (
	"noteforestx_server/utils"

	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

type ArticleService struct {
	utils utils.Utils
	Db    *gorm.DB
	Rdb   *redis.Client
}
