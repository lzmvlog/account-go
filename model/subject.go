package model

// Subject 科目实体
type Subject struct {
	Id int `gorm:"primary_key" json:"id"`
	// 科目名称
	SubName string `json:"subName"`
	// 父级id
	ParentId int `json:"parentId"`
	// 借贷方向 1 借 2 贷
	Direction int `json:"direction"`
	// 科目代码
	Code string `json:"code"`
	// 是否启用 1 启用 2 禁用
	IsEnable int `gorm:"force" json:"isEnable"`
	// 用户id
	UserId int `json:"userId"`
	// 用户昵称
	UserName string `json:"userName"`
}
