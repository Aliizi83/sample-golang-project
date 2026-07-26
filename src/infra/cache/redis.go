package cache

import (
	"context"
	"fmt"
	"time"

	"github.com/Aliizi83/sample-golang-project/src/config"
	"github.com/redis/go-redis/v9"
)

var redisClient *redis.Client

func InitRedis(cfg *config.Config) error {
	redisClient = redis.NewClient(&redis.Options{
		Addr:         fmt.Sprintf("%s:%s", cfg.Redis.Host, cfg.Redis.Port),
		Password:     cfg.Redis.Password,
		DialTimeout:  cfg.Redis.DialTimeout * time.Second,
		ReadTimeout:  cfg.Redis.ReadTimeout * time.Second,
		WriteTimeout: cfg.Redis.WriteTimeout * time.Second,
		PoolSize:     cfg.Redis.PoolSize,
		PoolTimeout:  cfg.Redis.PoolTimeout * time.Second,
	})

	ctx := context.Background()

	_, err := redisClient.Ping(ctx).Result()
	if err != nil {
		return err
	}
	return nil

}

func GetRedis() *redis.Client {
	return redisClient
}

func CloseRedis() {
	redisClient.Close()
}
