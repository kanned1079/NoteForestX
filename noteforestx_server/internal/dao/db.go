package dao

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"noteforestx_server/internal/config"
	"noteforestx_server/utils"
)

func (this *DaoInstance) InitDbConn() {
	this.logger.PrintInfo("create new db conn...")
	newDbConn, err := gorm.Open(mysql.New(mysql.Config{
		DSN: fmt.Sprintf("%s:%s@%s(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
			config.ExistingAppConfig.DbConfig.Username,
			config.ExistingAppConfig.DbConfig.Password,
			config.ExistingAppConfig.DbConfig.Protocol,
			config.ExistingAppConfig.DbConfig.Host,
			config.ExistingAppConfig.DbConfig.Port,
			config.ExistingAppConfig.DbConfig.Database),
		DefaultStringSize:         256,
		DisableDatetimePrecision:  true,
		DontSupportRenameIndex:    true,
		DontSupportRenameColumn:   true,
		SkipInitializeWithVersion: true,
	}), &gorm.Config{
		SkipDefaultTransaction: false,
		Logger:                 logger.Default.LogMode(logger.Warn),
	})
	if err != nil {
		panic(fmt.Sprintf("failed to open database: %v", err))
	}
	this.DbDao = newDbConn

	if this.DbDao.Exec(`SELECT 1 + 1;`).Error != nil {
		panic(fmt.Sprintf("failed to connect to database: %v", err))
	}
	utils.Logger{}.PrintSuccess("Server is connected and tested.")
}
