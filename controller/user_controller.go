package controller

import (
	"account-go/common"
	"account-go/model"
	"account-go/model/bo"
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
	user := bo.UserBo{}
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
	// 获取参数
	userBo := bo.UserBo{}
	c.ShouldBindJSON(&userBo)

	var user model.User
	err := db.Where("user_name = ?", userBo.UserName).First(&user).Error
	if err != nil {
		util.Response(c, http.StatusUnauthorized, 401, nil, err.Error())
		return
	}

	// 验证密码是否通过
	//if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(userBo.Password)); err != nil {
	//	util.Response(c, http.StatusUnprocessableEntity, 422, nil, "密码错误")
	//}

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
