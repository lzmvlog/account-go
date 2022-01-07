package util

// PageDetail 分页查询统一返回结构体
type PageDetail struct {
	// 数据列表
	DataList interface{} `json:"dataList"`
	// 当前页
	CurrentPage int `json:"current"`
	// 当前数量
	Count int `json:"count"`
	// 每页显示数量
	Size int `json:"size"`
}
