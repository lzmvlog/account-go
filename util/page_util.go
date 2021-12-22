package util

//type Pageable struct {
//	Page  int `json:"page"`
//	Size  int `json:"size"`
//	Total int `json:"total"`
//}

// PageDetail 分页查询统一返回结构体
type PageDetail struct {
	// 数据列表
	DataList interface{} `json:"dataList"`
	// 当前页
	CurrentPage string `json:"currentPage"`
	// 当前数量
	Count int `json:"count"`
	// 下一页
	NextPage bool `json:"nextPage"`
	// 每页显示数量
	PageSize string `json:"pageSize"`
}

func GetPage(p PageDetail) PageDetail {
	count, pageSize := p.CurrentPage, p.PageSize
	return PageDetail{
		DataList:    nil,
		CurrentPage: count,
		Count:       0,
		NextPage:    false,
		PageSize:    pageSize,
	}
}
