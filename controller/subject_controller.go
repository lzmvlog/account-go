package controller

import (
	"account-go/model"
	"account-go/util"
	"fmt"
	"github.com/gin-gonic/gin"
	"strconv"
)

// List 获取科目列表
func List(c *gin.Context) {
	var sub []model.Subject
	util.DB.Where("is_enable = 0").Find(&sub)
	util.Success(c, gin.H{"subject": sub}, "")
}

// Page 分页信息
func Page(c *gin.Context) {
	var sub []model.Subject
	page, size := c.Query("page"), c.Query("size")

	pageInt, err := strconv.Atoi(page)
	if err != nil {
		fmt.Println(nil)
	}

	sizeInt, err := strconv.Atoi(size)
	if err != nil {
		fmt.Println(nil)
	}
	var count int
	util.DB.Model(model.Subject{}).Where("is_enable = 0").Count(&count)

	// Limit 么也显示多少条 Offset 从第几条数据开始
	util.DB.Model(model.Subject{}).Where("is_enable = 0").Limit(sizeInt).Offset(pageInt - 1*sizeInt).Find(&sub)
	util.Success(c, gin.H{"page": util.PageDetail{DataList: sub, Count: count, CurrentPage: page, PageSize: size}}, "")
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
	c.ShouldBindJSON(&sub)
	util.DB.Model(model.Subject{}).Update(&sub)
	util.Success(c, gin.H{}, "")
}

// FindOne 根据id查询
func FindOne(c *gin.Context) {
	id := c.Query("id")

	idInt, err := strconv.Atoi(id)
	if err != nil {
		fmt.Println(nil)
	}

	sub := model.Subject{Id: idInt}
	util.DB.Model(model.Subject{}).Find(&sub)
	util.Success(c, gin.H{"subject": sub}, "")
}
