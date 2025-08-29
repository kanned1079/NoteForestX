package dao

import (
	"noteforestx_server/internal/models"
	"noteforestx_server/utils"
	"os"
)

const signFile = "./db_migrated_sign"

func (this *DaoInstance) AutoMigrateTables() {
	var logger utils.Logger
	logger.PrintInfo("start auto migrate tables")
	if _, err := os.Stat(signFile); err == nil {
		logger.PrintSuccess("skip auto migrate tables")
		return
	}
	if err := ExistingDbDaoInst.DbDao.AutoMigrate(&models.User{}, &models.Document{}); err != nil {
		logger.PrintError("auto migrate tables failed: ", err)
		return
	}
	if f, err := os.Create(signFile); err == nil {
		defer func() {
			if err := f.Close(); err != nil {
				logger.PrintError("close sign file failed: ", err)
				return
			}
		}()
	} else {
		logger.PrintError("create sign file failed: ", err)
		return
	}
	logger.PrintSuccess("auto migrate tables success")
}
