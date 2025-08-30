package dao

import (
	"fmt"
	"noteforestx_server/internal/models"
	"os"
	"time"
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
	if err := ExistingDbDaoInst.DbDao.AutoMigrate(
		&models.User{},
		&models.Document{},
		&models.IllustrationTag{},
		&models.Illustration{},
	); err != nil {
		this.logger.PrintError("auto migrate tables failed: ", err)
		os.Exit(1)
	}
	time.Sleep(time.Second * 2)
}

func (this *DaoInstance) AutoMigrateTables(force bool) {
	this.logger.DisplayLoadingAnime("running database migrations...")
	if force {
		this.logger.PrintWarn("force mode enabled, ignoring sign file...")
		this.runMigrate()
		this.createSignFile()
		fmt.Println()
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
	this.logger.StopLoadingAnime()
	this.logger.PrintSuccess("auto migrate tables success")
}
