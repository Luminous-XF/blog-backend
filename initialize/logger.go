package initialize

import (
	"blog-backend/global"
	"blog-backend/pkg/logger"
)

func initLogger() *logger.Logger {
	return logger.InitLogger(global.CONFIG)
}
