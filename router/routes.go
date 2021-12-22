package router

import (
	"account-go/controller"
	"github.com/gin-gonic/gin"
)

func CollectRouter(r *gin.Engine) *gin.Engine {
	sub := r.Group("/sub")
	{
		sub.GET("/list", controller.List)
		sub.POST("/save", controller.Save)
		sub.POST("/update", controller.Update)
		sub.POST("/page", controller.Page)
	}
	return r
}
