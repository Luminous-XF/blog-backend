// Package redis Redis工具包
package redis

import (
    "blog-backend/config"
    "blog-backend/pkg/logger"
    "context"
    "errors"
    redislib "github.com/redis/go-redis/v9"
    "sync"
    "time"
)

var (
    // rdb 包内单例变量
    rdb *RDB
    // once 确保 rdb 单例变量只被实例化一次
    once sync.Once
)

// RDB Redis 服务
type RDB struct {
    Client  *redislib.Client
    Context context.Context
}

// InitRedis 初始化 Redis 并返回一个 redis 实例
func InitRedis() *RDB {
    once.Do(func() {
        rdb = connect()
    })

    return rdb
}

// connect 连接 Redis
func connect() *RDB {
    cfg := &config.CONFIG.RedisConfig

    if len(cfg.Addr) == 0 {
        panic("redis addr is empty")
    }

    rds := &RDB{}
    rds.Context = context.Background()

    rds.Client = redislib.NewClient(&redislib.Options{
        Addr:     cfg.Addr,
        Password: cfg.Password,
        DB:       cfg.DB,
    })

    if err := rds.Ping(); err != nil {
        panic(err)
    }

    return rds
}

// Ping 测试 Redis 连接是否正常
func (rdb *RDB) Ping() error {
    _, err := rdb.Client.Ping(rdb.Context).Result()
    return err
}

// Set 存储 Key 对应的 value, 并设置 expiration 过期时间
func (rdb *RDB) Set(key string, value interface{}, expiration time.Duration) bool {
    if err := rdb.Client.Set(rdb.Context, key, value, expiration).Err(); err != nil {
        logger.LogIf(err)
        return false
    }

    return true
}

// Get 获取 key 对应的 value
func (rdb *RDB) Get(key string) string {
    res, err := rdb.Client.Get(rdb.Context, key).Result()
    if err != nil {
        if !errors.Is(err, redislib.Nil) {
            logger.ErrorString("Redis", "Get", err.Error())
        }
        return ""
    }

    return res
}

func (rdb *RDB) GetWithScan(key string, res interface{}) bool {
    if err := rdb.Client.Get(rdb.Context, key).Scan(res); err != nil {
        if !errors.Is(err, redislib.Nil) {
            logger.ErrorString("Redis", "GetWithScan", err.Error())
        }
    }

    return true
}

// Del 删除存储在 Redis 里的数据, 支持多个 key 传参
func (rdb *RDB) Del(keys ...string) bool {
    if err := rdb.Client.Del(rdb.Context, keys...).Err(); err != nil {
        logger.ErrorString("Redis", "Del", err.Error())
        return false
    }

    return true
}

// Exists 判断一个 key 是否存在, 内部错误和 redis.Nil 都返回 false
func (rdb *RDB) Exists(key string) bool {
    _, err := rdb.Client.Exists(rdb.Context, key).Result()
    if err != nil {
        if !errors.Is(err, redislib.Nil) {
            logger.ErrorString("Redis", "Exists", err.Error())
        }
        return false
    }

    return true
}

// FlushDB 清空当前 redis db 中所有数据
func (rdb *RDB) FlushDB() bool {
    if err := rdb.Client.FlushDB(rdb.Context).Err(); err != nil {
        logger.ErrorString("Redis", "FlushDB", err.Error())
        return false
    }

    return true
}

// Incr
// 一个参数时, 该参数作为 key, 将 key 对应的值加 1
// 两个参数时, 第一个参数为 key, 第二个参数为 value(int64), 将 key 对应的值加上 value
func (rdb *RDB) Incr(args ...interface{}) bool {
    switch len(args) {
    case 1:
        key := args[0].(string)
        if err := rdb.Client.Incr(rdb.Context, key).Err(); err != nil {
            logger.ErrorString("Redis", "Increment", err.Error())
            return false
        }
    case 2:
        key := args[0].(string)
        value := args[1].(int64)
        if err := rdb.Client.IncrBy(rdb.Context, key, value).Err(); err != nil {
            logger.ErrorString("Redis", "Increment", err.Error())
            return false
        }
    default:
        logger.ErrorString("Redis", "IncrBy", "Illegal parameters")
        return false
    }

    return true
}

// Decr
// 一个参数时, 该参数作为 key, 将 key 对应的值减 1
// 两个参数时, 第一个参数为 key, 第二个参数为 value(int64), 将 key 对应的值减去 value
func (rdb *RDB) Decr(args ...interface{}) bool {
    switch len(args) {
    case 1:
        key := args[0].(string)
        if err := rdb.Client.Decr(rdb.Context, key).Err(); err != nil {
            logger.ErrorString("Redis", "Decr", err.Error())
            return false
        }
    case 2:
        key := args[0].(string)
        value := args[1].(int64)
        if err := rdb.Client.DecrBy(rdb.Context, key, value).Err(); err != nil {
            logger.ErrorString("Redis", "Decr", err.Error())
            return false
        }
    default:
        logger.ErrorString("Redis", "Decr", "Illegal parameters")
        return false
    }

    return true
}
