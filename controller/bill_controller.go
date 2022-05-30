package controller

import (
	"account-go/model"
	"account-go/util"
	"github.com/gin-gonic/gin"
	"strconv"
)

// ListBill 获取科目列表
func ListBill(c *gin.Context) {
	var bills []model.Bill
	err := util.DB.Model(model.Bill{}).Find(&bills).Error
	if err != nil {
		util.Fail(c, err.Error())
		return
	}
	util.Success(c, gin.H{"bill": bills}, "")
}

// PageBill 分页信息
func PageBill(c *gin.Context) {

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

	var count int64
	var user model.User
	user = GetUser(c)

	errOne := util.DB.Model(model.Bill{}).Where("user_id = ?", user.Id).Count(&count).Error
	if errOne != nil {
		util.Fail(c, err.Error())
		return
	}

	list := make([]model.Bill, 0)
	// Limit 么也显示多少条 Offset 从第几条数据开始
	errFind := util.DB.Model(model.Bill{}).Where("user_id = ?", user.Id).Limit(sizeInt).Offset((pageInt - 1) * sizeInt).Find(&list).Error
	if errFind != nil {
		util.Fail(c, err.Error())
		return
	}

	util.Success(c, gin.H{"page": util.PageDetail{DataList: list, Count: count, CurrentPage: pageInt, Size: sizeInt}}, "")
}

// SaveBill 保存账单表
func SaveBill(c *gin.Context) {
	tx := util.DB.Begin()
	bill := model.Bill{}

	c.ShouldBindJSON(&bill)

	bill = GetBillSubject(bill, c)
	// 这里需要注意 create 传入的是结构体的指针
	err := tx.Model(model.Bill{}).Create(&bill).Error
	if err != nil {
		util.Fail(c, err.Error())
		tx.Rollback()
		return
	}
	tx.Commit()
	util.Success(c, gin.H{}, "新增成功")
}

// UpdateBill 更新科目表
func UpdateBill(c *gin.Context) {
	tx := util.DB.Begin()
	bill := model.Bill{}

	c.ShouldBindJSON(&bill)

	bill = GetBillSubject(bill, c)
	err := tx.Model(model.Bill{}).Where(bill.Id).Updates(&bill).Error
	if err != nil {
		util.Fail(c, "update bill error")
		tx.Rollback()
		return
	}
	tx.Commit()
	util.Success(c, gin.H{}, "")
}

// GetBillSubject 获取账单得科目信息
func GetBillSubject(bill model.Bill, c *gin.Context) model.Bill {
	var user model.User
	// 获取用户信息
	user = GetUser(c)
	bill.UserId = user.Id
	bill.UserName = user.UserName

	var sub model.Subject
	sub = SelectSubjectOne(bill.SubId)

	bill.SubName = sub.SubName
	bill.Direction = sub.Direction
	return bill
}

// FindBillOne 根据id查询
func FindBillOne(c *gin.Context) {
	id := c.Query("id")

	idInt, err := strconv.Atoi(id)
	if err != nil {
		util.Fail(c, err.Error())
		return
	}

	bill := model.Bill{Id: idInt}
	util.DB.Model(model.Bill{}).Find(&bill)
	util.Success(c, gin.H{"bill": bill}, "")
}
