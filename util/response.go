package util

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// Response 返回统一信息
func Response(c *gin.Context, httpStatus int, code int, data gin.H, message string) {
	c.JSON(httpStatus, gin.H{"code": code, "data": data, "msg": message})
}

// Success 请求成功
func Success(c *gin.Context, data gin.H, message string) {
	if message == "" {
		message = "success"
	}
	Response(c, http.StatusOK, 200, data, message)
}

// FailMessage 请求失败
func FailMessage(c *gin.Context, data gin.H, message string) {
	Response(c, http.StatusOK, 500, data, message)
}

// Fail 请求失败
func Fail(c *gin.Context, message string) {
	if message == "" {
		message = "程序异常！"
	}
	Response(c, http.StatusOK, 500, gin.H{}, message)
}
