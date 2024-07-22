package redis_mq

import (
    "blog-backend/pkg/logger"
    "context"
    "encoding/json"
    "fmt"
    "github.com/redis/go-redis/v9"
    "go.uber.org/zap"
    "strconv"
    "time"
)

// handlerFunc 用于处理消息的函数类型
type handlerFunc func(msg Message)

// 默认处理函数, 打印消息
func defaultHandler(msg Message) {
    logger.Info("RedisMQ:", zap.String("Message", fmt.Sprintf("%+v", msg)))
}

// consumer 消费者相关字段和方法
type consumer struct {
    ctx      context.Context
    duration time.Duration
    ch       chan []string
    handler  handlerFunc
}

// newConsumer 创建一个消费者实例
func newConsumer(ctx context.Context, handler handlerFunc) *consumer {
    return &consumer{
        ctx:      ctx,
        duration: time.Second,
        ch:       make(chan []string, 1000),
        handler:  handler,
    }
}

func (c *consumer) listen(rdb *redis.Client, topic string) {
    // 从哈希表中获取数据并处理
    go func() {
        for {
            select {
            case ret := <-c.ch:
                // 从哈希表中批量获取数据信息
                key := topic + HashSuffix
                res, err := rdb.HMGet(c.ctx, key, ret...).Result()
                if err != nil {
                    logger.Error("RedisMQ", zap.String("HMGet error", err.Error()))
                }

                if len(res) > 0 {
                    rdb.HDel(c.ctx, key, ret...)
                }

                msg := Message{}
                for _, v := range res {
                    // 由于哈希表和有序集合操作不是原子操作
                    // 可能会出现删除了集合中的数据但哈希表中数据未删除的情况
                    if v == nil {
                        continue
                    }

                    str := v.(string)
                    _ = json.Unmarshal([]byte(str), &msg)

                    // 处理逻辑
                    go c.handler(msg)
                }
            }
        }
    }()

    // 定时器用于定时获取消息并处理
    ticker := time.NewTicker(c.duration)
    defer ticker.Stop()
    for {
        select {
        case <-c.ctx.Done(): // 上下文取消, 退出监听
            logger.Info("RedisMQ", zap.String("Consumer quit", c.ctx.Err().Error()))
            return
        case <-ticker.C: // 定时获取消息
            // 从 Redis 中读取消息
            minValue := strconv.Itoa(0)
            maxValue := strconv.Itoa(int(time.Now().Unix()))
            opt := &redis.ZRangeBy{
                Min: minValue,
                Max: maxValue,
            }

            key := topic + SetSuffix
            res, err := rdb.ZRangeByScore(c.ctx, key, opt).Result()
            if err != nil {
                logger.Error("RedisMQ", zap.String("ZRangeByScore error", err.Error()))
                return
            }

            // 获取到数据
            if len(res) > 0 {
                // 从有序集合中移除数据
                rdb.ZRemRangeByScore(c.ctx, key, minValue, maxValue)

                // 写入通道, 进行哈希表处理
                c.ch <- res
            }
        }
    }
}
