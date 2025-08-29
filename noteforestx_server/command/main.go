package main

import (
	"github.com/gin-gonic/gin"
	"noteforestx_server/internal/config"
	"noteforestx_server/internal/dao"
	"noteforestx_server/internal/router"
	"noteforestx_server/utils"
	"sync"
)

func init() {
	// setup app
	utils.ShowFigure("NoteForestX", "0.0.1", "Kanna")

	config.ExistingAppConfig.ReadConfigFile("./config/config.yaml")
	dao.ExistingDbDaoInst = dao.NewDbInstance(1)
	dao.ExistingDbDaoInst.AutoMigrateTables()
}

func main() {
	var logger utils.Logger
	logger.PrintInfo("all done. start NoteForestX server")

	var existingRouterInst = router.NewRouterInstance(1, gin.DebugMode)

	var wg sync.WaitGroup

	wg.Add(1)
	go func() {
		defer wg.Done()
		existingRouterInst.StartApiServices()
	}()

	wg.Wait()
}
