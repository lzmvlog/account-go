package model

import "time"

// User 用户信息
type User struct {
	Id int `gorm:"primary_key" json:"id"`
	// 用户昵称
	UserName string `json:"userName"`
	// 用户密码
	Password string `json:"password"`
	// 创建时间
	CreateDate time.Time `json:"createDate"`
	// 是否启用 0 启用 1 禁用
	IsEnable int `json:"isEnable"`
}
