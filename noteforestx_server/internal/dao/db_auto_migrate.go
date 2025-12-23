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

//func (this *DaoInstance) runMigrate() {
//	if err := ExistingDbDaoInst.DbDao.AutoMigrate(
//		&models.User{},
//		&models.Document{},
//		&models.IllustrationImage{},
//		&models.IllustrationTag{},
//		&models.IllustrationAuthor{},
//		&models.Illustration{},
//		&models.IllustrationTagMapping{},
//	); err != nil {
//		this.logger.PrintError("auto migrate tables failed: ", err)
//		os.Exit(1)
//	}
//	time.Sleep(time.Second * 2)
//}

func (this *DaoInstance) dropAllTables() {
	// 按依赖顺序删除，避免外键约束错误
	tables := []struct {
		TableName string
		Model     interface{}
	}{
		{"illustration_images", &models.IllustrationImage{}},
		{"illustration_tag_mapping", &models.IllustrationTagMapping{}},
		{"illustrations", &models.Illustration{}},
		{"illustration_tags", &models.IllustrationTag{}},
		{"illustration_authors", &models.IllustrationAuthor{}},
		{"users", &models.User{}},
		{"documents", &models.Document{}},
		{"article", &models.Article{}},
		{"article_tag", &models.ArticleTag{}},
	}

	for _, t := range tables {
		if err := this.DbDao.Migrator().DropTable(t.Model); err != nil {
			this.logger.PrintError("drop table "+t.TableName+" failed: ", err)
			os.Exit(1)
		}
		this.logger.PrintInfo("dropped table: " + t.TableName)
	}

	this.logger.PrintInfo("all tables dropped successfully")
}

func (this *DaoInstance) runMigrate() {
	if err := ExistingDbDaoInst.DbDao.AutoMigrate(
		&models.User{},
		&models.Document{},
		&models.IllustrationAuthor{},
		&models.IllustrationTag{},
		&models.Illustration{},           // Illustration 必须在前
		&models.IllustrationImage{},      // 依赖 Illustration
		&models.IllustrationTagMapping{}, // 依赖 Illustration & Tag
		&models.Article{},
		&models.ArticleTag{},
	); err != nil {
		this.logger.PrintError("auto migrate tables failed: ", err)
		os.Exit(1)
	}
	time.Sleep(time.Second * 2)
}

func (this *DaoInstance) showTimeout(seconds int) {
	for i := seconds; i > 0; i-- {
		this.logger.PrintWarn(fmt.Sprintf("tables will be removed in %d s...", i))
		time.Sleep(time.Second)
	}
}

func (this *DaoInstance) AutoMigrateTables(force bool) {
	if force {
		this.logger.PrintWarn("force mode enabled, ignoring sign file...")
		this.logger.PrintError("all tables will be removed in 10s, you can press \"^C\" to cancel...")
		this.showTimeout(10)
		//time.Sleep(time.Second * 10)
		//this.dropAllTables()
		this.logger.DisplayLoadingAnime("running database migrations (force)...")
		this.runMigrate()
		this.createSignFile()
		fmt.Println()
		this.logger.PrintSuccess("auto migrate tables success (forced)")
		return
	}

	this.logger.DisplayLoadingAnime("running database migrations...")
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
