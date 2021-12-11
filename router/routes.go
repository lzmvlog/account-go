package router

import (
	"account-go/controller"
	"github.com/gin-gonic/gin"
)

func CollectRouter(r *gin.Engine) *gin.Engine {
	r.GET("/sub/list", controller.List)
	r.POST("/sub/save", controller.Save)
	return r
}
