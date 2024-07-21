package initialize

import (
	"blog-backend/pkg/redis"
)

func initRedis() *redis.RDB {
    return redis.InitRedis()
}
