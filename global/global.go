package global

import (
	"blog-backend/config"
	"blog-backend/pkg/redis"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

var (
	CONFIG *config.Config
	GDB    *gorm.DB
	RDB    *redis.RDB
	Engine *gin.Engine
)
