package dao

import (
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
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
	daoInst.InitRedisConn()

	return daoInst
}
