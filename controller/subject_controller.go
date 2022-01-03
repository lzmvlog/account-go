package controller

import (
	"account-go/model"
	"account-go/util"
	"github.com/gin-gonic/gin"
	"strconv"
)

// ListSubject 获取科目列表
func ListSubject(c *gin.Context) {
	var sub []model.Subject
	err := util.DB.Model(model.Subject{}).Where("is_enable = 0").Find(&sub).Error
	if err != nil {
		util.Fail(c, err.Error())
		return
	}
	util.Success(c, gin.H{"subject": sub}, "")
}

// PageSubject 分页信息
func PageSubject(c *gin.Context) {
	var sub []model.Subject
	page, size := c.Query("page"), c.Query("size")

	pageInt, err := strconv.Atoi(page)
	if err != nil {
		util.Fail(c, err.Error())
		return
	}

	sizeInt, err := strconv.Atoi(size)
	if err != nil {
		util.Fail(c, err.Error())
		return
	}

	var count int
	errOne := util.DB.Model(model.Subject{}).Where("is_enable = 0").Count(&count).Error
	if errOne != nil {
		util.Fail(c, err.Error())
		return
	}

	// Limit 么也显示多少条 Offset 从第几条数据开始
	errFind := util.DB.Model(model.Subject{}).Where("is_enable = 0").Limit(sizeInt).Offset(pageInt - 1*sizeInt).Find(&sub).Error
	if errFind != nil {
		util.Fail(c, err.Error())
		return
	}

	util.Success(c, gin.H{"page": util.PageDetail{DataList: sub, Count: count, CurrentPage: page, PageSize: size}}, "")
}

// SaveSubject 保存科目表表
func SaveSubject(c *gin.Context) {
	tx := util.DB.Begin()
	sub := model.Subject{}
	c.ShouldBindJSON(&sub)
	// 这里需要注意 create 传入的是结构体的指针
	err := tx.Model(model.Subject{}).Create(&sub).Error
	if err != nil {
		util.Fail(c, err.Error())
		tx.Rollback()
		return
	}
	tx.Commit()
	util.Success(c, gin.H{}, "")
}

// UpdateSubject 更新科目表
func UpdateSubject(c *gin.Context) {
	tx := util.DB.Begin()
	sub := model.Subject{}
	c.ShouldBindJSON(&sub)
	err := tx.Model(model.Subject{}).Where(sub.Id).Updates(sub).Error
	if err != nil {
		util.Fail(c, err.Error())
		tx.Rollback()
		return
	}
	tx.Commit()
	util.Success(c, gin.H{}, "")
}

// FindSubjectOne  根据id查询
func FindSubjectOne(c *gin.Context) {
	id := c.Query("id")

	idInt, err := strconv.Atoi(id)
	if err != nil {
		util.Fail(c, err.Error())
		return
	}

	sub := model.Subject{Id: idInt}
	util.DB.Model(model.Subject{}).Find(&sub)
	util.Success(c, gin.H{"subject": sub}, "")
}
