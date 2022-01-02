package model

// User 账单信息
type User struct {
	Id int `gorm:"primary_key" json:"id"`
	// 用户昵称
	UserName string `json:"userName"`
	// 用户密码
	Password string `json:"password"`
}
