package initialize

import (
	"blog-backend/global"
	"blog-backend/pkg/redis"
	redislib "github.com/redis/go-redis/v9"
)

func initRedis() *redislib.Client {
	return redis.InitRedis(&global.CONFIG.RedisConfig)
}
