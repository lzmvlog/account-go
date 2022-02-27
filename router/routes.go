package router

import (
	"account-go/controller"
	"github.com/gin-gonic/gin"
)

// CollectRouter 路由设置
func CollectRouter(r *gin.Engine) *gin.Engine {
	// 用户路由
	user := r.Group("/api")
	{
		user.POST("/login", controller.Login)
		user.POST("/register", controller.Register)
		user.GET("/pageUser", controller.PageUser)
	}

	// 科目路由
	sub := r.Group("/sub")
	{
		sub.GET("/listSub", controller.ListSubject)
		sub.POST("/saveSub", controller.SaveSubject)
		sub.POST("/updateSub", controller.UpdateSubject)
		sub.GET("/pageSub", controller.PageSubject)
		sub.GET("/findOneSub", controller.FindSubjectOne)
		sub.GET("/disable/:id", controller.Disable)
	}

	// 账单路由
	bill := r.Group("/bill")
	{
		bill.GET("/listBill", controller.ListBill)
		bill.POST("/saveBill", controller.SaveBill)
		bill.GET("/pageBill", controller.PageBill)
		bill.GET("/findOneBill", controller.FindBillOne)
	}
	return r
}
