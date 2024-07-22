// Package redis_mq 消费实体定义
package redis_mq

import (
    "encoding/json"
    "github.com/google/uuid"
    "time"
)

// Message 定义消息结构
type Message struct {
    ID          string      `json:"id"`
    CreateTime  time.Time   `json:"createTime"`
    ConsumeTime time.Time   `json:"consumeTime"`
    Body        interface{} `json:"body"`
}

// NewMessage 用于创建消息实体
func NewMessage(id string, consumeTime time.Time, body interface{}) *Message {
    if len(id) == 0 {
        id = uuid.New().String()
    }

    return &Message{
        ID:          id,
        CreateTime:  time.Now(),
        ConsumeTime: consumeTime,
        Body:        body,
    }
}

// GetScore 用于返回消息的分数
func (msg *Message) GetScore() float64 {
    return float64(msg.ConsumeTime.Unix())
}

// GetID 用于返回消息的 ID
func (msg *Message) GetID() string {
    return msg.ID
}

// MarshalBinary 将消息结构体序列化为二进制数据
func (msg *Message) MarshalBinary() ([]byte, error) {
    return json.Marshal(msg)
}

// UnmarshalBinary 用于将二进制数据反序列化为消息结构体
func (msg *Message) UnmarshalBinary(data []byte) error {
    return json.Unmarshal(data, msg)
}
