package controller

import (
	"account-go/model"
	"account-go/util"
	"github.com/gin-gonic/gin"
)

// List 获取科目列表
func List(c *gin.Context) {
	db := util.GetDB()
	var sub []model.Subject
	db.Where("is_enable = 0").Find(&sub)
	util.Success(c, gin.H{"subject": sub}, "")
}

// Save 保存科目表表
func Save(c *gin.Context) {
	db := util.GetDB()
	var sub model.Subject
	c.BindJSON(&sub)
	err := db.Save(sub)
	if err != nil {
		util.Fail(c, err.Error.Error())
	}
}
