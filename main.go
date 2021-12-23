package main

import (
	"account-go/config"
	"account-go/middlewares"
	"account-go/router"
	"account-go/util"
	"github.com/gin-gonic/gin"
)

func main() {

	// 初始化数据库
	db := util.InitDB()
	// 关闭数据库连接
	defer db.Close()
	r := gin.Default()
	r.Use(middlewares.Cors())
	r = router.CollectRouter(r)
	port := config.Config.Server.Port
	if port != "" {
		panic(r.Run(":" + port))
	}
	r.Run(config.Config.Server.Port) // listen and serve on 0.0.0.0:8080
}
