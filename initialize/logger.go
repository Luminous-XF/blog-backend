package initialize

import (
	"blog-backend/global"
	"blog-backend/pkg/logger"
)

func initLogger() {
	logger.InitLogger(global.CONFIG)
}
