package main

import (
	"os"

	"github.com/premwitthawas/basic-api/config"
	"github.com/premwitthawas/basic-api/modules/servers"
	"github.com/premwitthawas/basic-api/pkg/database"
)

func envPath() string {
	if len(os.Args) == 1 {
		return ".env"
	} else {
		return os.Args[1]
	}
}

func main() {
	cfg := config.LoadConfig(envPath())
	db := database.DbConnect(cfg.Db())
	defer db.Close()
	servers.CreateServer(cfg, db).Start()
}
