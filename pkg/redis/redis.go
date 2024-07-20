package redis

import (
	"blog-backend/config"
	"context"
	redislib "github.com/redis/go-redis/v9"
	"sync"
)

var (
	rdb  *redislib.Client
	once sync.Once
)

func InitRedis(cfg *config.RedisConfig) *redislib.Client {
	once.Do(func() {
		rdb = connect(cfg)
	})
	return rdb
}

func connect(cfg *config.RedisConfig) *redislib.Client {
	if len(cfg.Addr) == 0 {
		panic("redis addr is empty")
	}

	rdb = redislib.NewClient(&redislib.Options{
		Addr:     cfg.Addr,
		Password: cfg.Password,
		DB:       cfg.DB,
	})

	if _, err := rdb.Ping(context.Background()).Result(); err != nil {
		panic(err)
	}

	return rdb
}
