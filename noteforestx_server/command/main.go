package main

import (
	"fmt"
	"noteforestx_server/internal/config"
	"noteforestx_server/internal/dao"
	"noteforestx_server/internal/router"
	"noteforestx_server/utils"
	"os"
	"sync"
)

const (
	name    = "NoteForestX"
	version = "0.0.1"
)

func main() {
	fmt.Printf("%s v%s\n", name, version)

	if len(os.Args) < 2 {
		printHelp()
		os.Exit(1)
	}

	cmd := os.Args[1]

	switch cmd {
	case "server", "-s":
		runServer()
	case "init", "-i":
		runInit()
	case "reset":
		runReset()
	case "help", "-h", "--help":
		printHelp()
	default:
		fmt.Printf("Unknown command: %s", cmd)
		printHelp()
		os.Exit(1)
	}
}

func printHelp() {
	fmt.Print(`
Usage: go run main.go [command] [options]

Available commands:
  server              Run the NoteForestX server
  init [--force]      Initialize (e.g. migrate database tables at first time)
                      --force   Force re-initialize (e.g. skip sign file check and force migrate)
  reset               Reset the admin password
  help                Show this help message

`)
}

func runInit() {
	var logger utils.Logger
	config.ExistingAppConfig.ReadConfigFile("./config/config.yaml")
	dao.ExistingDbDaoInst = dao.NewDbInstance(1)

	args := os.Args
	force := false

	if len(args) > 2 {
		for _, arg := range args[2:] {
			if arg == "--force" {
				force = true
				break
			} else {
				logger.PrintError("unknown command: ", arg)
				os.Exit(1)
			}
		}
	}

	if force {
		logger.PrintWarn("force init, will drop sign file and force migrate tables")
		dao.ExistingDbDaoInst.AutoMigrateTables(true)
	} else {
		dao.ExistingDbDaoInst.AutoMigrateTables(false)
	}
}

func runReset() {

}

func runServer() {
	// Load config first
	config.ExistingAppConfig.ReadConfigFile("./config/config.yaml")
	dao.ExistingDbDaoInst = dao.NewDbInstance(1)
	var logger utils.Logger
	logger.PrintInfo("All done. Starting NoteForestX server...")

	// Init router
	existingRouterInst := router.NewRouterInstance(1, config.ExistingAppConfig.Runtime.Mode)

	var wg sync.WaitGroup
	wg.Add(1)

	go func() {
		defer wg.Done()
		existingRouterInst.RegisterApiServices()
		existingRouterInst.StartAndServe(config.ExistingAppConfig.Runtime.ListeningPort)
	}()

	wg.Wait()
}
