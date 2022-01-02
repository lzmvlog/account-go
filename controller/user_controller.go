package controller

import (
	"account-go/common"
	"account-go/model"
	"account-go/model/dto"
	"account-go/util"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
	"net/http"
)

// Register 注册/开通账号
func Register(c *gin.Context) {
	db := util.DB

	// 获取参数
	name := c.PostForm("name")
	phone := c.PostForm("phone")
	password := c.PostForm("password")

	user := model.User{
		UserName: name,
		Password: password,
	}
	// 数据验证
	if !userDataValidation(user, c) {
		return
	}
	// 判断手机号是否存在
	if isPhoneExits(db, phone) {
		util.Response(c, http.StatusUnprocessableEntity, 442, nil, "当前手机号已经存在")
		return
	}

	// 如果名称没有传，给一个随机的字符串
	if len(name) == 0 {
		util.Response(c, http.StatusUnprocessableEntity, 442, nil, "用户名称不能为空")
		return
	}

	// 加密密码
	bcryptPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		util.Response(c, http.StatusUnprocessableEntity, 500, nil, "加密出错")
		return
	}

	// 新建学生信息
	user.Password = string(bcryptPassword)
	db.Create(&user)

	util.Success(c, nil, "注册成功")
}

// isPhoneExits 判断电话是否存在
func isPhoneExits(db *gorm.DB, phone string) bool {
	var user model.User
	db.Where("phone = ?", phone).First(&user)
	if user.Id != 0 {
		return true
	}
	return false
}

// Login 登录方法
func Login(c *gin.Context) {
	db := util.DB
	userName, password := c.Query("userName"), c.Query("password")

	if !userDataValidation(model.User{
		Password: password,
	}, c) {
		return
	}

	var user model.User
	db.Where("user_name = ?", userName).First(&user)
	if user.Id == 0 {
		util.Response(c, http.StatusUnprocessableEntity, 422, nil, "用户不存在")
		return
	}

	// 验证密码是否通过
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		util.Response(c, http.StatusUnprocessableEntity, 422, nil, "密码错误")
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
