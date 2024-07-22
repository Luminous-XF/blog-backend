package redis_mq

import (
    "context"
    "github.com/redis/go-redis/v9"
    "sync"
)

const (
    HashSuffix = ":hash" // Redis 键后缀, 用于哈希表
    SetSuffix  = ":set"  // Redis 键后后缀, 用于集合
)

var (
    once sync.Once // 用于实现单例
)

type Queue struct {
    ctx context.Context // 上下文

    // redis
    rdb   *redis.Client // Redis 客户端
    topic string        // 主题

    producer *producer // 生产者
    consumer *consumer // 消费者
}

func NewQueue(ctx context.Context, rdb *redis.Client, opts ...Option) *Queue {
    var q *Queue

    once.Do(func() {
        // 定义默认的选项
        defaultOptions := Options{
            topic:   "topic",
            handler: defaultHandler,
        }

        for _, apply := range opts {
            apply(&defaultOptions)
        }

        // 创建 Queue 实例
        q = &Queue{
            ctx:      ctx,
            rdb:      rdb,
            topic:    defaultOptions.topic,
            producer: newProducer(ctx),
            consumer: newConsumer(ctx, defaultOptions.handler),
        }
    })

    return q
}

func (q *Queue) Start() {
    // 启动消费者的监听
    go q.consumer.listen(q.rdb, q.topic)
}

func (q *Queue) Publish(msg *Message) (int64, error) {
    // 发布消息
    return q.producer.publish(q.rdb, q.topic, msg)
}
