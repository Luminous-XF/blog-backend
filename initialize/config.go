package initialize

import (
	"blog-backend/config"
	"blog-backend/pkg/viper"
)

func initConfig() *config.Config {
	return viper.InitConfig()
}
