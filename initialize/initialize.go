package initialize

import (
	"blog-backend/global"
)

func InitProject() {
	// 读取配置文件
	global.CONFIG = initConfig()

	// 初始化日志工具
	global.Logger = initLogger()

	// 连接 MySQL 数据库
	global.GDB = initDB()

	// 连接 Redis 数据库
	global.RDB = initRedis()

	// 启动 Gin
	global.Engine = initGin()

	// 初始化参数校验工具
	initValidator()
}
