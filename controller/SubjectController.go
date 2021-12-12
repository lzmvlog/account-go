package controller

import (
	"account-go/model"
	"account-go/util"
	"github.com/gin-gonic/gin"
)

// List 获取科目列表
func List(c *gin.Context) {
	var sub []model.Subject
	util.DB.Where("is_enable = 0").Find(&sub)
	util.Success(c, gin.H{"subject": sub}, "")
}

// Save 保存科目表表
func Save(c *gin.Context) {
	sub := model.Subject{}
	c.BindJSON(&sub)
	err := util.DB.Create(sub)
	if err != nil {
		util.Fail(c, "save subject error")
		return
	}
	util.Success(c, gin.H{}, "")
}
