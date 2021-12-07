package util

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"github.com/spf13/viper"
)

var DB *gorm.DB

// InitDB 初始化数据库连接
func InitDB() *gorm.DB {
	driverName := viper.GetString("datasource.driverName")
	host := viper.GetString("datasource.host")
	port := viper.GetString("datasource.port")
	database := viper.GetString("datasource.database")
	userName := viper.GetString("datasource.username")
	password := viper.GetString("datasource.password")
	charset := viper.GetString("datasource.charset")
	args := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=true",
		userName, password, host, port, database, charset)

	// 官方依赖 ：gorm.io/gorm 工具包 ：github.com/jinzhu/gorm
	// 官方的 gorm.Open() 方法： Open(dialector Dialector, opts ...Option)
	// 工具包方法: Open(dialect string, args ...interface{})
	db, err := gorm.Open(driverName, args)
	if err != nil {
		panic("failed to connect database,err:" + err.Error())
	}

	// 自动创建数据表
	// db.AutoMigrate(&Student{})
	DB = db
	return db
}

// GetDB 获取数据库连接
func GetDB() *gorm.DB {
	return DB
}
