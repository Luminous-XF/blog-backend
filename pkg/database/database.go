package database

import (
	"blog-backend/config"
	"blog-backend/pkg/logger"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"sync"
)

var (
	db   *gorm.DB
	once sync.Once
)

func connect(dbCfg *config.DatabaseConfig, mysqlCfg *config.MySQLConfig) {
	if len(dbCfg.Name) == 0 {
		panic("database name is empty")
	}

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		dbCfg.User,
		dbCfg.Password,
		dbCfg.Host,
		dbCfg.Port,
		dbCfg.Name,
	)

	var err error
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger:      logger.NewGormLogger(),
		PrepareStmt: true,
	})

	if err != nil {
		panic(fmt.Errorf("connect to mysql use dataSourceName %s failed: %s", dsn, err))
	}

	// 设置数据库连接池参数, 提高并发性能
	sqlDB, _ := db.DB()
	// 设置数据库连接池最大连接数
	sqlDB.SetMaxOpenConns(mysqlCfg.MaxOpenConnections)
	// 设置连接池最大允许的空闲连接shu, 若没有 sql 任务需要执行的连接数大于 20, 超过的连接会被连接池关闭
	sqlDB.SetMaxIdleConns(mysqlCfg.MaxIdleConnections)
}

func InitDB(cfg *config.DatabaseConfig, mysqlCfg *config.MySQLConfig) *gorm.DB {
	once.Do(func() {
		connect(cfg, mysqlCfg)
	})
	return db
}
