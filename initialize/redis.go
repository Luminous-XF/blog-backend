package initialize

import (
	"blog-backend/global"
	"blog-backend/pkg/redis"
)

func initRedis() *redis.RDB {
	return redis.InitRedis(&global.CONFIG.RedisConfig)
}
