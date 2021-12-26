package model

import (
	"github.com/shopspring/decimal"
	"time"
)

// Bill 账单信息
type Bill struct {
	Id int `gorm:"primary_key" json:"id"`
	// 科目id
	SubId int `json:"subId"`
	// 科目名称
	SubName string `json:"subName"`
	// 借贷方向 0 借 1 贷
	Direction int `json:"direction"`
	// 金额数量
	Amount decimal.Decimal `json:"amount"`
	// 创建日志
	CreateDate time.Time `json:"createDate"`
	// 备注
	Remark string `json:"remark"`
}
