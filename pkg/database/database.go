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
    // db 包内单例变量
    db *gorm.DB
    // once 确保 db 单例变量只被实例化一次
    once sync.Once
)

// InitDB 初始化数据库, 并返回一个 gorm.DB 实例
func InitDB() *gorm.DB {
    once.Do(func() {
        db = connect()
    })

    return db
}

// connect 建立数据库连接
func connect() *gorm.DB {
    dbCfg := &config.CONFIG.DatabaseConfig
    if len(dbCfg.Name) == 0 {
        panic("database name is empty")
    }

    // 配置数据库连接串
    dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
        dbCfg.User,
        dbCfg.Password,
        dbCfg.Host,
        dbCfg.Port,
        dbCfg.Name,
    )

    // 连接 MySQL
    var err error
    dbs, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
        Logger:      logger.NewGormLogger(),
        PrepareStmt: true,
    })

    if err != nil {
        panic(fmt.Errorf("connect to mysql use dataSourceName %s failed: %s", dsn, err))
    }

    // 设置数据库连接池参数, 提高并发性能
    mysqlCfg := &config.CONFIG.MySQLConfig
    MysqlDB, _ := dbs.DB()
    // 设置数据库连接池最大连接数
    MysqlDB.SetMaxOpenConns(mysqlCfg.MaxOpenConnections)
    // 设置连接池最大允许的空闲连接shu, 若没有 sql 任务需要执行的连接数大于 20, 超过的连接会被连接池关闭
    MysqlDB.SetMaxIdleConns(mysqlCfg.MaxIdleConnections)

    return dbs
}
