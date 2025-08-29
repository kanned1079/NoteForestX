package dao

import (
	"context"
	"fmt"
	"github.com/redis/go-redis/v9"
	"noteforestx_server/internal/config"
	"time"
)

func (this *DaoInstance) InitRedisConn() {
	this.logger.PrintInfo("create new redis conn...")

	this.RdbDao = redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%d", config.ExistingAppConfig.RedisConfig.Host, config.ExistingAppConfig.RedisConfig.Port),
		Password: config.ExistingAppConfig.RedisConfig.Password,
		DB:       config.ExistingAppConfig.RedisConfig.Database,
	})

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := this.RdbDao.Ping(ctx).Err(); err != nil {
		this.logger.PrintError("failed to connect redis: %v", err)
		panic(err)
	}

	this.logger.PrintSuccess("redis connection established successfully")
}
