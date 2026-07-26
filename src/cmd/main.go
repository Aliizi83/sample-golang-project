package main

import (
	"log"

	"github.com/Aliizi83/sample-golang-project/src/api"
	"github.com/Aliizi83/sample-golang-project/src/config"
	"github.com/Aliizi83/sample-golang-project/src/infra/cache"
)

func main() {

	cfg := config.GetConfig()
	err := cache.InitRedis(cfg)
	defer cache.CloseRedis()
	if err != nil {
		log.Fatal("Redis connection failed: " + err.Error())
		return
	}

	api.InitServer(cfg)
}
