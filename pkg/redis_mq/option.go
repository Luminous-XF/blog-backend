package redis_mq

type Option func(*Options)

type Options struct {
    topic   string
    handler handlerFunc
}

// WithTopic 用于设置选项中的 topic 字段
func WithTopic(topic string) Option {
    return func(opts *Options) {
        opts.topic = topic
    }
}

// WithHandler 用于设置选项中的 handler 字段
func WithHandler(handler handlerFunc) Option {
    return func(opts *Options) {
        opts.handler = handler
    }
}
