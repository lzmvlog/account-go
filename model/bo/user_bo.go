package bo

// UserBo 用户信息
type UserBo struct {
	// 用户昵称
	UserName string `json:"userName"`
	// 用户密码
	Password string `json:"password"`
}
