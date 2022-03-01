package controller

import (
	"account-go/common"
	"account-go/model"
	"account-go/model/bo"
	"account-go/model/dto"
	"account-go/util"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"net/http"
	"strconv"
	"time"
)

// Register 注册/开通账号
func Register(c *gin.Context) {
	db := util.DB

	// 获取参数
	user := model.User{}
	c.ShouldBindJSON(&user)

	userInfo := model.User{
		UserName: user.UserName,
		Password: user.Password,
	}

	// 数据验证
	if !userDataValidation(userInfo, c) {
		return
	}

	// 如果名称没有传，给一个随机的字符串
	if len(user.UserName) == 0 {
		util.Response(c, http.StatusUnprocessableEntity, 442, nil, "用户名称不能为空")
		return
	}

	// 加密密码
	bcryptPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		util.Response(c, http.StatusUnprocessableEntity, 500, nil, "加密出错")
		return
	}

	// 添加时间
	user.CreateDate = time.Now()

	// 初始化用户信息
	user.Password = string(bcryptPassword)
	errc := db.Create(&user).Error
	if errc != nil {
		util.Fail(c, errc.Error())
		return
	}

	util.Success(c, nil, "注册成功")
}

// Login 登录方法
func Login(c *gin.Context) {
	db := util.DB
	// 获取参数
	userBo := bo.UserBo{}
	c.ShouldBindJSON(&userBo)

	var user model.User
	err := db.Where("user_name = ?  and is_enable = 0", userBo.UserName).First(&user).Error
	if err != nil {
		util.Response(c, http.StatusUnauthorized, 401, nil, err.Error())
		return
	}

	// 验证密码是否通过
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(userBo.Password)); err != nil {
		util.Response(c, http.StatusUnprocessableEntity, 422, nil, "密码错误")
		return
	}

	if user.Id == 0 {
		util.Response(c, http.StatusUnauthorized, 401, nil, "当前账号未启用")
		return
	}

	// 发放token
	token, err := common.ReleaseToken(user)
	if err != nil {
		util.Response(c, http.StatusUnprocessableEntity, 500, nil, "系统异常")
	}

	util.Success(c, gin.H{"token": token}, "登录成功")

}

// userDataValidation 用户数据校验
func userDataValidation(user model.User, c *gin.Context) bool {

	if len(user.Password) < 6 {
		util.Response(c, http.StatusUnprocessableEntity, 422, nil, "密码不能小于6位")
		return false
	}

	return true
}

// Info 获取用户信息
func Info(c *gin.Context) {
	user, _ := c.Get("user")
	util.Success(c, gin.H{"user": dto.ToUserDTO(user.(model.User))}, "")
}

// GetUser 获取
func GetUser(c *gin.Context) model.User {
	value, exists := c.Get("user")
	if !exists {
		util.Fail(c, "请重新登录")
	}
	data, err := json.Marshal(value)
	if err != nil {
		util.Fail(c, "序列化异常")
	}
	var user model.User

	err = json.Unmarshal(data, &user)
	if err != nil {
		util.Fail(c, "序列化异常")
	}
	return user
}

// PageUser 分页信息
func PageUser(c *gin.Context) {
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
	errOne := util.DB.Model(model.User{}).Count(&count).Error
	if errOne != nil {
		util.Fail(c, err.Error())
		return
	}

	var user []model.User
	// Limit 么也显示多少条 Offset 从第几条数据开始
	errFind := util.DB.Model(model.User{}).Limit(sizeInt).Offset((pageInt - 1) * sizeInt).Find(&user).Error
	if errFind != nil {
		util.Fail(c, err.Error())
		return
	}

	util.Success(c, gin.H{"page": util.PageDetail{DataList: user, Count: count, CurrentPage: pageInt, Size: sizeInt}}, "")
}

// DisableUser 禁用科目
func DisableUser(c *gin.Context) {
	id := c.Param("id")
	var user model.User
	tx := util.DB.Begin()
	err := tx.Model(model.User{}).Where("id= ?", id).Find(&user).Error
	if err != nil {
		util.Fail(c, err.Error())
		return
	}

	if user.IsEnable == 0 {
		user.IsEnable = 1
	} else {
		user.IsEnable = 0
	}

	err = tx.Model(&user).UpdateColumn("is_enable", user.IsEnable).Error
	if err != nil {
		util.Fail(c, err.Error())
		tx.Rollback()
		return
	}
	tx.Commit()

	util.Success(c, gin.H{}, "")
}

// GetOneUser 获取一个用户的信息
func GetOneUser(c *gin.Context) {
	id := c.Param("id")

	idInt, err := strconv.Atoi(id)
	if err != nil {
		util.Fail(c, err.Error())
		return
	}

	user := model.User{Id: idInt}
	util.DB.Model(model.User{}).Find(&user)
	util.Success(c, gin.H{"user": user}, "")
}

// UpdateUser 更新用户
func UpdateUser(c *gin.Context) {
	tx := util.DB.Begin()
	user := model.User{}
	c.ShouldBindJSON(&user)
	err := tx.Model(model.User{}).Where(user.Id).Updates(user).Error
	if err != nil {
		util.Fail(c, err.Error())
		tx.Rollback()
		return
	}
	tx.Commit()
	util.Success(c, gin.H{}, "")
}
