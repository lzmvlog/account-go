package main

import (
	"account-go/router"
	"account-go/util"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"os"
)

func main() {
	InitConfig()
	// 初始化数据库
	db := util.InitDB()
	// 关闭数据库连接
	defer db.Close()
	r := gin.Default()
	r = router.CollectRouter(r)
	port := viper.GetString("server.port")
	if port != "" {
		panic(r.Run(":" + port))
	}
	r.Run(":8090") // listen and serve on 0.0.0.0:8080
}

// InitConfig 初始化配置
func InitConfig() {
	workDir, _ := os.Getwd()
	viper.SetConfigName("application")
	viper.SetConfigType("yml")
	viper.AddConfigPath(workDir + "/config")
	err := viper.ReadInConfig()
	if err != nil {
		panic("read config error")
	}
}
