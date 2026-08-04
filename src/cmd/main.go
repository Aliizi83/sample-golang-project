package main

import (
	"github.com/Aliizi83/sample-golang-project/src/api"
	"github.com/Aliizi83/sample-golang-project/src/config"
	"github.com/Aliizi83/sample-golang-project/src/infra/cache"
	"github.com/Aliizi83/sample-golang-project/src/infra/presistence/db"
	"github.com/Aliizi83/sample-golang-project/src/infra/presistence/migrations"
	"github.com/Aliizi83/sample-golang-project/src/pkg/logging"
)

// @securityDefinitions.apikey AuthBearer
// @in header
// @name Authorization
func main() {

	cfg := config.GetConfig()
	err := cache.InitRedis(cfg)
	logger := logging.NewLogger(cfg)
	defer cache.CloseRedis()
	if err != nil {
		logger.Fatal(err, logging.Redis, logging.Startup, err.Error(), nil)
		return
	}

	if err := db.InitDB(cfg); err != nil {
		logger.Fatal(err, logging.Postgres, logging.Startup, err.Error(), nil)
	}

	migrations.UpP_1()
	defer db.CloseDB()

	api.InitServer(cfg)
}
