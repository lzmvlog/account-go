package util

import (
	"account-go/config"
	_ "github.com/go-sql-driver/mysql"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"os"
	"time"
)

var DB *gorm.DB

// InitDB 初始化数据库连接
func InitDB() *gorm.DB {
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer（日志输出的目标，前缀和日志包含的内容——译者注）
		logger.Config{
			SlowThreshold:             time.Second, // 慢 SQL 阈值
			LogLevel:                  logger.Info, // 日志级别
			IgnoreRecordNotFoundError: true,        // 忽略ErrRecordNotFound（记录未找到）错误
			Colorful:                  true,        // 禁用彩色打印
		})
	db, err := gorm.Open(mysql.Open(config.Config.Datasource.Url), &gorm.Config{
		Logger: newLogger,
	})
	if err != nil {
		panic("failed to connect database,err:" + err.Error())
	}

	// 自动创建数据表
	// db.AutoMigrate(&Student{})
	DB = db
	return db
}
