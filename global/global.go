package global

import (
	"blog-backend/config"
	"blog-backend/pkg/logger"
	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

var (
	CONFIG *config.Config
	GDB    *gorm.DB
	RDB    *redis.Client
	Logger *logger.Logger
	Engine *gin.Engine
)
