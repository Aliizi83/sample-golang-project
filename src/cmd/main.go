package main

import (
	"github.com/Aliizi83/sample-golang-project/src/api"
	"github.com/Aliizi83/sample-golang-project/src/config"
	"github.com/Aliizi83/sample-golang-project/src/infra/cache"
)

func main() {

	cfg := config.GetConfig()
	cache.InitRedis(cfg)
	defer cache.CloseRedis()

	api.InitServer(cfg)
}
