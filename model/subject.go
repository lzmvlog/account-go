package model

// Subject 科目实体
type Subject struct {
	Id int `gorm:"primary_key" json:"id"`
	// 科目名称
	SubName string `json:"subName"`
	// 父级id
	ParentId int `json:"parentId"`
	// 借贷方向 0 借 1 贷
	Direction int `json:"direction"`
	// 科目代码
	Code string `json:"code"`
	// 是否启用 0 启用 1 禁用
	IsEnable int `gorm:"force" json:"isEnable"`
}
