package redis_mq

import (
    "context"
    "github.com/redis/go-redis/v9"
)

// producer 生产者结构体
type producer struct {
    ctx context.Context
}

// newProducer 创建一个生产者实例
func newProducer(ctx context.Context) *producer {
    return &producer{ctx: ctx}
}

// publish 发布消息到 Redis 中
func (p *producer) publish(rdb *redis.Client, topic string, msg *Message) (int64, error) {
    z := redis.Z{
        Member: msg.GetID(),
        Score:  msg.GetScore(),
    }

    // 将消息写入有序集合
    setKey := topic + SetSuffix
    n, err := rdb.ZAdd(p.ctx, setKey, z).Result()
    if err != nil {
        return n, err
    }

    // 将消息写入哈希表
    hashKey := topic + HashSuffix
    return rdb.HSet(p.ctx, hashKey, msg.GetID(), msg).Result()
}
