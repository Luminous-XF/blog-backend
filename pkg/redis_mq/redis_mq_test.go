package redis_mq

import (
    "context"
    "fmt"
    "github.com/google/uuid"
    "github.com/redis/go-redis/v9"
    "testing"
    "time"
)

func TestRedisMQ(t *testing.T) {
    rdb := redis.NewClient(&redis.Options{
        Addr:     "127.0.0.1:6379",
        Password: "",
        DB:       1,
    })
    ctx := context.Background()

    _, err := rdb.Ping(ctx).Result()
    if err != nil {
        t.Error(err)
        t.Fail()
    }

    redisMQ := NewQueue(
        context.Background(), rdb,
        WithTopic("Send-Message"),
        WithHandler(func(msg Message) {
            fmt.Printf("%#v\n", msg)
        }))

    redisMQ.Start()

    id := uuid.New().String()
    logMsg := NewMessage(id, time.Now(), map[string]interface{}{
        "id":     123,
        "action": "Login",
    })

    v, err := redisMQ.Publish(logMsg)
    if err != nil {
        t.Error(err)
        t.Fail()
    }
    fmt.Println(v)

    time.Sleep(3 * time.Second)
}
