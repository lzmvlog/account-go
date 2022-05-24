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
	util.InitDB()
	r := gin.Default()
	// 中间件 处理跨域
	r.Use(middlewares.Cors())
	// 处理认证信息
	r.Use(middlewares.AuthMiddleware())
	r = router.CollectRouter(r)
	port := config.Config.Server.Port
	if port != "" {
		panic(r.Run(":" + port))
	}
	r.Run(config.Config.Server.Port) // listen and serve on 0.0.0.0:8080
}
