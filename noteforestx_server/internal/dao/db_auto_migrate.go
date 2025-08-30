package dao

import (
	"noteforestx_server/internal/models"
	"os"
)

const signFile = "./db_migrated_sign"

func (this *DaoInstance) createSignFile() {
	f, err := os.Create(signFile)
	if err != nil {
		this.logger.PrintError("create sign file failed: ", err)
		return
	}
	defer func() {
		if err := f.Close(); err != nil {
			this.logger.PrintError("close sign file failed: ", err)
		}
	}()
}

func (this *DaoInstance) runMigrate() {
	this.logger.PrintInfo("running database migrations...")
	if err := ExistingDbDaoInst.DbDao.AutoMigrate(
		&models.User{},
		&models.Document{},
	); err != nil {
		this.logger.PrintError("auto migrate tables failed: ", err)
		return
	}
}

func (this *DaoInstance) AutoMigrateTables(force bool) {
	if force {
		this.logger.PrintWarn("force mode enabled, ignoring sign file...")
		this.runMigrate()
		this.createSignFile()
		this.logger.PrintSuccess("auto migrate tables success (forced)")
		return
	}

	this.logger.PrintInfo("start auto migrate tables")
	if _, err := os.Stat(signFile); err == nil {
		this.logger.PrintWarn("skip auto migrate tables, sign file exists (use --force to override)")
		return
	}

	this.runMigrate()
	this.createSignFile()
	this.logger.PrintSuccess("auto migrate tables success")
}
