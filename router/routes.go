package router

import (
	"account-go/controller"
	"github.com/gin-gonic/gin"
)

// CollectRouter 路由设置
func CollectRouter(r *gin.Engine) *gin.Engine {
	sub := r.Group("/sub")
	{
		sub.GET("/list", controller.List)
		sub.POST("/save", controller.Save)
		sub.POST("/update", controller.Update)
		sub.GET("/page", controller.Page)
		sub.GET("/findOne", controller.FindOne)
	}
	return r
}
