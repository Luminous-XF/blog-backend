package initialize

import (
	"blog-backend/pkg/database"
	"gorm.io/gorm"
)

func initDB() *gorm.DB {
    return database.InitDB()
}
