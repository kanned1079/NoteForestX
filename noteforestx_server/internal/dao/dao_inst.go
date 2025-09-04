package dao

import (
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
	"noteforestx_server/internal/config"
	"noteforestx_server/utils"
)

type DaoInstance struct {
	Id     int32
	DbDao  *gorm.DB
	RdbDao *redis.Client
	logger utils.Logger
}

var ExistingDbDaoInst *DaoInstance

func NewDbInstance(id int32) *DaoInstance {
	var daoInst = &DaoInstance{Id: id}
	daoInst.InitDbConn()
	if config.ExistingAppConfig.RedisConfig.Enabled {
		daoInst.InitRedisConn()
	} else {
		daoInst.logger.PrintWarn("redis is disabled, init skipped. you can edit \"config.yaml\" to enable it.")
	}

	return daoInst
}
