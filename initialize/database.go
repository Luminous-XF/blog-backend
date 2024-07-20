package initialize

import (
	"blog-backend/global"
	"blog-backend/pkg/database"
	"gorm.io/gorm"
)

func initDB() *gorm.DB {
	return database.InitDB(&global.CONFIG.DatabaseConfig, &global.CONFIG.MySQLConfig)
}
