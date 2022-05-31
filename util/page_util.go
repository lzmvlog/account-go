package util

// PageDetail 分页查询统一返回结构体
type PageDetail struct {
	// 数据列表
	Records interface{} `json:"records"`
	// 当前页
	CurrentPage int `json:"current"`
	// 当前数量
	Total int64 `json:"total"`
	// 每页显示数量
	Size int `json:"size"`
}
