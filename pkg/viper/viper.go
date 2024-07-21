package viper

import (
    "blog-backend/config"
    "errors"
    "fmt"
    "github.com/fsnotify/fsnotify"
    viperlib "github.com/spf13/viper"
)

var vp *viperlib.Viper

func init() {
    vp = viperlib.New()
}

func InitConfig() *config.Config {
    return readConfig()
}

func readConfig() *config.Config {
    vp.AddConfigPath("config")
    vp.SetConfigName("config")
    vp.SetConfigType("yaml")

    if err := vp.ReadInConfig(); err != nil {
        var configFileNotFoundError viperlib.ConfigFileNotFoundError
        if !errors.As(err, &configFileNotFoundError) {
            panic(fmt.Errorf("file not found: %s", err))
        } else {
            panic(fmt.Errorf("failed to read file: %s", err))
        }
    }

    // 解析配置文件
    var cfg config.Config
    if err := vp.Unmarshal(&cfg); err != nil {
        panic(fmt.Errorf("failed to parse configuration file: %s", err))
    }

    // 监听配置文件
    vp.WatchConfig()
    vp.OnConfigChange(func(e fsnotify.Event) {
        if err := vp.Unmarshal(&config.CONFIG); err != nil {
            panic(fmt.Errorf("failed to parse configuration file: %s", err))
        }
    })

    return &cfg
}
