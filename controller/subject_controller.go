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
	err := util.DB.Model(model.Subject{}).Where("is_enable = 1").Find(&sub).Error
	if err != nil {
		util.FailMessage(c, err.Error())
		return
	}
	util.Success(c, sub)
}

// PageSubject 分页信息
func PageSubject(c *gin.Context) {
	page, size := c.Query("page"), c.Query("size")

	subName := c.Query("subName")

	pageInt, err := strconv.Atoi(page)
	if err != nil {
		util.FailMessage(c, err.Error())
		return
	}

	sizeInt, err := strconv.Atoi(size)
	if err != nil {
		util.FailMessage(c, err.Error())
		return
	}

	var count int64
	errOne := util.DB.Model(model.Subject{}).Count(&count).Error
	if errOne != nil {
		util.FailMessage(c, err.Error())
		return
	}

	list := make([]model.Subject, 0)
	// Limit 么也显示多少条 Offset 从第几条数据开始
	if subName != "" {
		err = util.DB.Model(model.Subject{}).Where("sub_name LIKE ?", "%"+subName+"%").Limit(sizeInt).Offset((pageInt - 1) * sizeInt).Find(&list).Count(&count).Error
	} else {
		err = util.DB.Model(model.Subject{}).Limit(sizeInt).Offset((pageInt - 1) * sizeInt).Find(&list).Error
	}

	if err != nil {
		util.FailMessage(c, err.Error())
		return
	}

	util.PageSuccess(c, list, count, pageInt, sizeInt)
}

// SaveSubject 保存科目表表
func SaveSubject(c *gin.Context) {
	tx := util.DB.Begin()
	sub := model.Subject{}
	c.ShouldBindJSON(&sub)

	var user model.User
	user = GetUser(c)

	sub.UserId = user.Id
	sub.UserName = user.UserName

	// 这里需要注意 create 传入的是结构体的指针
	err := tx.Model(model.Subject{}).Create(&sub).Error
	if err != nil {
		util.FailMessage(c, err.Error())
		tx.Rollback()
		return
	}
	tx.Commit()
	util.Success(c, nil)
}

// UpdateSubject 更新科目表
func UpdateSubject(c *gin.Context) {
	tx := util.DB.Begin()
	sub := model.Subject{}
	c.ShouldBindJSON(&sub)
	err := tx.Model(model.Subject{}).Where(sub.Id).Updates(sub).Error
	if err != nil {
		util.FailMessage(c, err.Error())
		tx.Rollback()
		return
	}
	tx.Commit()
	util.Success(c, nil)
}

// FindSubjectOne  根据id查询
func FindSubjectOne(c *gin.Context) {
	id := c.Query("id")

	idInt, err := strconv.Atoi(id)
	if err != nil {
		util.FailMessage(c, err.Error())
		return
	}
	var sub model.Subject
	sub = SelectSubjectOne(idInt)

	util.Success(c, sub)
}

// SelectSubjectOne 根据借贷方向获取科目信息
func SelectSubjectOne(subId int) model.Subject {
	sub := model.Subject{Id: subId}
	util.DB.Model(model.Subject{}).Find(&sub)
	return sub
}

// DisableSubject 禁用科目
func DisableSubject(c *gin.Context) {
	id := c.Param("id")
	var sub model.Subject
	tx := util.DB.Begin()
	err := tx.Model(model.Subject{}).Where("id = ?", id).Find(&sub).Error
	if err != nil {
		util.FailMessage(c, err.Error())
		return
	}

	if sub.IsEnable == 1 {
		sub.IsEnable = 2
	} else {
		sub.IsEnable = 1
	}

	err = tx.Model(&sub).UpdateColumn("is_enable", sub.IsEnable).Error
	if err != nil {
		util.FailMessage(c, err.Error())
		tx.Rollback()
		return
	}

	tx.Commit()
	util.Success(c, nil)
}
