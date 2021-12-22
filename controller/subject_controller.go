package controller

import (
	"account-go/model"
	"account-go/util"
	"fmt"
	"github.com/gin-gonic/gin"
)

// List 获取科目列表
func List(c *gin.Context) {
	var sub []model.Subject
	util.DB.Where("is_enable = 0").Find(&sub)
	util.Success(c, gin.H{"subject": sub}, "")
}

func Page(c *gin.Context) {
	var sub []model.Subject
	page, size := c.Query("page"), c.Query("size")
	util.GetPage(util.PageDetail{
		DataList:    nil,
		CurrentPage: c.Query("page"),
		Count:       0,
		NextPage:    false,
		PageSize:    c.Query("size"),
	})
	// Limit 么也显示多少条 Offset 第几页
	util.DB.Where("is_enable = 0").Limit(size).Offset(page).Find(&sub)
	//util.DB.Where()
	util.Success(c, gin.H{"subject": sub}, "")
}

// Save 保存科目表表
func Save(c *gin.Context) {
	sub := model.Subject{}
	c.ShouldBindJSON(&sub)
	err := util.DB.Create(sub)
	if err != nil {
		util.Fail(c, "save subject error")
		return
	}
	util.Success(c, gin.H{}, "")
}

// Update 更新科目表
func Update(c *gin.Context) {
	sub := model.Subject{}
	c.BindJSON(&sub)
	fmt.Println(sub)
}
