package main

import (
	"log"
	"noteforestx_server/internal/config"
	"noteforestx_server/internal/dao"
)

func init() {
	// setup app
	config.ExistingAppConfig.ReadConfigFile("./config/config.yaml")
	dao.ExistingDbDaoInst = dao.NewDbInstance(1)
}

func main() {

	log.Println("Start server")

}
