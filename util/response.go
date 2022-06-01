package util

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

const (
	SUCCESS = "SUCCESS"
	FAIL    = "FAIL"
)

// Response 返回统一信息
func Response(c *gin.Context, httpStatus int, message string) {
	c.JSON(httpStatus, gin.H{"code": httpStatus, "msg": message})
}

// ResponseFail 返回统一信息
func ResponseFail(c *gin.Context, httpStatus int, message string) {
	c.JSON(httpStatus, gin.H{"code": httpStatus, "msg": message})
}

// ResponseSuccess 返回统一信息
func ResponseSuccess(c *gin.Context, httpStatus int, data interface{}, message string) {
	c.JSON(httpStatus, gin.H{"code": httpStatus, "data": data, "msg": message})
}

// ResponsePage 分页返回统一信息
func ResponsePage(c *gin.Context, httpStatus int, data interface{}, message string) {
	c.JSON(httpStatus, gin.H{"code": httpStatus, "data": data, "msg": message})
}

// PageSuccess 请求成功
func PageSuccess(c *gin.Context, data interface{}, total int64, page int, size int) {
	var pageInfo PageDetail
	pageInfo.Records = data
	pageInfo.CurrentPage = page
	pageInfo.Size = size
	pageInfo.Total = total
	ResponsePage(c, http.StatusOK, pageInfo, SUCCESS)
}

// Success 请求成功
func Success(c *gin.Context, data interface{}) {
	ResponseSuccess(c, http.StatusOK, data, SUCCESS)
}

// FailMessage 请求失败
func FailMessage(c *gin.Context, message string) {
	ResponseFail(c, http.StatusInternalServerError, message)
}

// Fail 请求失败
func Fail(c *gin.Context) {
	Response(c, http.StatusInternalServerError, FAIL)
}
